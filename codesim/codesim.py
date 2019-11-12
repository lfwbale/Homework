from difflib import SequenceMatcher
import sys
import pygments.token
import pygments.lexers
def tokenize(filename):
    """对代码进行预处理，并返回处理后的元素列表"""
    file = open(filename,'r')
    text = file.read()
    file.close()
    lexer = pygments.lexers.guess_lexer_for_filename(filename, text)
    tokens = lexer.get_tokens(text)
    tokens = list(tokens)
    count1=0
    count2=0
    #print(tokens)
    result = []
    lenT = len(tokens)
    for i in range(lenT):
        #将变量名都转换成N
        if tokens[i][0] == pygments.token.Name :
            result.append(('N', count1, count2))  #all variable names as 'N'
            count2 += 1
        #将字符串都转换为S
        elif tokens[i][0] in pygments.token.Literal.String:
            result.append(('S', count1, count2))  #all strings as 'S'
            count2 += 1
        #将函数名都转换成F
        elif tokens[i][0] in pygments.token.Name.Function:
            result.append(('F', count1, count2))   #user defined function names as 'F'
            count2 += 1
        #忽略括号 换行符 空格等干扰因素
        elif tokens[i][0] == pygments.token.Text or tokens[i][0] in pygments.token.Comment or tokens[i][0] in pygments.token.Punctuation:
            pass   #whitespaces and comments ignored
        else:
            result.append((tokens[i][1], count1, count2))
            #tuples in result-(each element e.g 'def', its position in original code file, position in cleaned up code/text)
            count2 += len(tokens[i][1])
        count1 += len(tokens[i][1])

    return result

def toString(arr):
    """将列表中每个元素重新连接成一个新的字符串"""
    String=''
    for i in arr:
        String += str(i[0])
    return String

#file = input("name:")
#syr=tokenize(file)
#print(syr)
#print(toString(syr))

#file1 = input("Enter the name of file1: ")
#file2 = input("Enter the name of file2: ")
file1 = sys.argv[1]
file2 = sys.argv[2]
tokens1 = tokenize(file1)
tokens2 = tokenize(file2)
Content1 = toString(tokens1)
Content2 = toString(tokens2)
SM = SequenceMatcher(None, Content1, Content2)
similarity_ratio = SM.ratio()
print(similarity_ratio)

