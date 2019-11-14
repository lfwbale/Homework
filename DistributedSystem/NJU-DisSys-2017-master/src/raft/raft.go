package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//


import "sync"
import "labrpc"
import "fmt"
import "time"
import "math/rand"	//用于生成随机数


// import "bytes"
// import "encoding/gob"


//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make().
type ApplyMsg struct {
	Index       int
	Command     interface{}
	UseSnapshot bool   // ignore for lab2; only used in lab3
	Snapshot    []byte // ignore for lab2; only used in lab3
}

type Log struct{
	term int 
}

const follower = 1
const candidate =2
const leader = 3


//
// A Go object implementing a single Raft peer.
//

type Raft struct {
	mu        sync.Mutex
	peers     []*labrpc.ClientEnd
	persister *Persister
	me        int // index into peers[]

	// Your data here.
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.
	currentTerm int   // latest term server has seen 当前任期
	votedFor    int   // candidatedId that received vote in current term 投票给某个candidateId
	status      int   // follower or candidate or leader _server当前的状态
	voteCount   int   // 当前candidate获得的选票数目
	log         []Log // log entries;each entry contains command 日志集合
					  // for state machine,and term when entry was received by leader
	heartbeatCh chan bool
	voteCh      chan bool
	leaderCh    chan bool
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {
	//获取节点当前状态
	var term int
	var isleader bool
	// Your code here.
	term = rf.currentTerm
	if rf.status == leader{
		isleader = true
	} else {
		isleader = false
	}
	return term, isleader
}

//
// save Raft's persistent state to stable storage,
// where it can later be retrieved after a crash and restart.
// see paper's Figure 2 for a description of what should be persistent.
//

func (rf *Raft) persist() {
	// Your code here.
	// Example:
	// w := new(bytes.Buffer)
	// e := gob.NewEncoder(w)
	// e.Encode(rf.xxx)
	// e.Encode(rf.yyy)
	// data := w.Bytes()
	// rf.persister.SaveRaftState(data)
}

//
// restore previously persisted state.
//

func (rf *Raft) readPersist(data []byte) {
	// Your code here.
	// Example:
	// r := bytes.NewBuffer(data)
	// d := gob.NewDecoder(r)
	// d.Decode(&rf.xxx)
	// d.Decode(&rf.yyy)
}

//
// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//

func (rf *Raft) Start(command interface{}) (int, int, bool) {
	//start agreement on a new log entry
	index := -1
	term := -1
	isLeader := true
	return index, term, isLeader

}

//
// the tester calls Kill() when a Raft instance won't
// be needed again. you are not required to do anything
// in Kill(), but it might be convenient to (for example)
// turn off debug output from this instance.
//
func (rf *Raft) Kill() {
	// Your code here, if desired.
}

//设置timeout 时间为150ms-300ms
func genElectionTimeout() time.Duration {
	return time.Duration(rand.Intn(150)+150) * time.Millisecond
}

//发出心跳
type AppendEntriesArgs struct {
	Term     int	//laeder's term
	LeaderId int	//so follower can redirect clients
}
//心跳回应
type AppendEntriesReply struct {
	Term    int		//currentTerm,for leader to update itself
	Success bool	//true if follower contained entry matching 
					//preLogIndex and preLogTerm
}
//要求投票
type RequestVoteArgs struct {
	// Your data here.
	Term         int	//candidate's term
	CandidateId  int	//candidate requesting vote
	LastLogIndex int	//index of candidate's last log entry
	LastLogTerm  int	//term of candidate's last log entry
}
//投票回应
type RequestVoteReply struct {
	// Your data here.
	Term        int		//currentTerm,for candidate to update itself
	VoteGranted bool	//true means candidate to update itself
	ServerId    int		//follower's id
}

func (rf *Raft) AppendEntries(args AppendEntriesArgs, reply *AppendEntriesReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	if args.Term < rf.currentTerm {
		// 当前 server 的 term 大于 leader 的 term , 
		// 更新 reply 中的 term, 让 leader 知道自己的 term 过期了
		reply.Term = rf.currentTerm
		reply.Success = false
		return
	} else if args.Term > rf.currentTerm {
		//leader不变 造反失败
		// 更新当前 server 的 term 并且将状态改变为 follower, 重置选举计时器
		rf.currentTerm = args.Term
		rf.status = follower
		rf.votedFor = -1
		reply.Term = args.Term
		reply.Success = true
	} else {
		reply.Success = true
	}
	go func() {
		rf.heartbeatCh <- true
	}()
}

func (rf *Raft) sendAppendEntries(server int, args AppendEntriesArgs, reply *AppendEntriesReply) bool {
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	rf.mu.Lock()
	defer rf.mu.Unlock()
	if ok {
		// 只有 leader 才进行转化
		if rf.status != leader || args.Term != rf.currentTerm {
			return ok
		}
		// 如果 reply 结果中的 Term 大于 Leader 节点的当前 Term，则 Leader 节点转换为 Follwer 状态
		if reply.Term > rf.currentTerm {
			rf.currentTerm = reply.Term
			rf.status = follower
			rf.votedFor = -1
			rf.persist()
			return ok
		}
	}
	return ok
}

func (rf *Raft) broadcastAppendEntries() {
	rf.mu.Lock()
	args := AppendEntriesArgs{rf.currentTerm, rf.me}
	defer rf.mu.Unlock()
	for i := 0; i < len(rf.peers); i++ {
		if rf.me == i || rf.status != leader {
			continue
		}
		go func(i int, args AppendEntriesArgs) {
			var reply AppendEntriesReply
			rf.sendAppendEntries(i, args, &reply)
		}(i, args)
	}
}

func (rf *Raft) RequestVote(args RequestVoteArgs, reply *RequestVoteReply) {
	// Your code here.
	rf.mu.Lock()
	defer rf.mu.Unlock()
	// fmt.Println("request:", rf.me, ", args-id:", args.CandidateId)
	if args.Term < rf.currentTerm {
		// 当前 server 的 term 大于 candidate 的 term , 拒绝投票, 更新 reply 中的 term
		//造反失败
		reply.Term = rf.currentTerm
		reply.VoteGranted = false
		return
	} else if args.Term > rf.currentTerm {
		// 当前 server 的 term 小于 candidate 的 term 
		// 更新自身 term 并且成为 follower , 投票给candidate
		reply.VoteGranted = true
		reply.ServerId = rf.me	
		rf.currentTerm = args.Term
		rf.status = follower
		rf.votedFor = args.CandidateId	//用于记票
		go func() { rf.voteCh <- true }()
	} else if rf.votedFor == -1 {
		//选出第一任leader
		// 当两者的term相等时, server为初始状态
		reply.VoteGranted = true
		reply.ServerId = rf.me
		reply.Term = rf.currentTerm
		rf.votedFor = args.CandidateId
		go func() { rf.voteCh <- true }()
	}
}


func (rf *Raft) sendRequestVote(server int, args RequestVoteArgs, reply *RequestVoteReply) bool {
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
		// fmt.Println(args.CandidateId, rf.me, reply.ServerId)
		rf.mu.Lock()
		defer rf.mu.Unlock()
		if ok {
			// 只有 candidate 才可以转换为 leader
			if rf.status != candidate || args.Term != rf.currentTerm {
				return ok
			}
	
			// reply 结果中的 Term 比 candidate 的 Term 大，则 candidate 转换为 follower 状态。
			if reply.Term > rf.currentTerm {
				rf.currentTerm = reply.Term
				rf.status = follower
				rf.votedFor = -1
				rf.persist()
			}
	
			// fmt.Println("reply term", reply.Term, "rf term", rf.currentTerm)
			// 反馈结果中有投票, 则当前 candidate 的 获得票+1, 并且判断活动票是否超过一半
			// 如果超过一半, 则则转换为 leader 状态
			if reply.VoteGranted {
				rf.voteCount += 1
				if rf.voteCount >= len(rf.peers)/2+1 {
					// fmt.Println("candidate", rf.me, "became a leader")
					rf.status = leader
					rf.votedFor = -1
					rf.leaderCh <- true
				}
			}
		}
	return ok
}

func (rf *Raft) broadcastRequestVote() {
	rf.mu.Lock()
	args := RequestVoteArgs{rf.currentTerm, rf.me, -1, 0}
	rf.mu.Unlock()
	// fmt.Println("-------------")
	// fmt.Println("candidate", rf.me)
	for i := 0; i < len(rf.peers); i++ {
		if rf.me == i || rf.status != candidate {
			continue
		}

		go func(i int) {
			var reply RequestVoteReply
			rf.sendRequestVote(i, args, &reply)
			// fmt.Println("server", i, reply.VoteGranted)
		}(i)
	}
}

func (rf *Raft) getLeader() {
	rf.mu.Lock()
	//开始下一任选举并给自己投上一票
	rf.currentTerm++
	rf.votedFor = rf.me
	rf.voteCount = 1
	rf.mu.Unlock()
	go rf.broadcastRequestVote()
	select {
	//超时 什么也不做
	case <-time.After(genElectionTimeout()):
	//收到心跳 造反失败 变成follower
	case <-rf.heartbeatCh:
		rf.status = follower
	//选举成功 成为leader
	case <-rf.leaderCh:
		rf.status = leader
		//rf.print()
	}
}

func (rf *Raft) startElection() {
	for {
		switch rf.status {
		case follower:
			select {
			case <-rf.heartbeatCh:
			case <-rf.voteCh:
			//没收到心跳和投票通知，超时，变成候选人
			case <-time.After(genElectionTimeout()):
				rf.status = candidate
			}
		case candidate:
			//发起投票
			rf.getLeader()
		case leader:
			//行使leader权利
			go rf.broadcastAppendEntries()
			time.Sleep(time.Duration(50) * time.Millisecond)
		}

	}
}

func Make(peers []*labrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me

	// Your initialization code here.
	rf.status = follower
	rf.votedFor = -1
	rf.currentTerm = 0
	rf.heartbeatCh = make(chan bool) //建立心跳通道
	rf.voteCh = make(chan bool)		 //建立投票通道
	rf.leaderCh = make(chan bool)	 //建立通知成为leader通道

	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())

	go rf.startElection()

	return rf
}

func (rf *Raft) print() {
	time.Sleep(time.Duration(rand.Intn(100)+150) * time.Millisecond)
	fmt.Print("Raft:", rf.me, "  status:", rf.status, "  term:", rf.currentTerm, "  voteFor:", rf.voteCount, " \n")
}