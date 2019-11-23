package com.test.jacoco;

import javax.imageio.event.IIOReadWarningListener;
import java.util.ArrayList;
import java.util.List;


public class HelloWorld {
    public HelloWorld() {

    }

    public String Method1() {
        return "Hello World";
    }

    public int Method2(int a, int b) {
        return a + b;
    }

    public int Method3(int a, int b, int c) {
        if ((a > 5 && b < 0) || c > 0) {
            System.out.println("Condition 1");
        } else if (a < 5 && c < -2) {
            System.out.println("Condition 2");
        } else {
            System.out.println("Condition 3");
        }
        return 0;
    }

    public int Method4(int a, int b, int c, int d, float e) {
        if (a == 0) {
            return 0;
        }
        int x = 0;
	    if ((a == b) || ((c == d) && (bug(a)))) {
        //if (((c == d) && (bug(a))) || (a == b)) {
            x = 1;
        }
        e = 1 / x;
        return 0;
    }

    public boolean bug(int a) {
        if (a == 5) return true;
        return false;
    }


    public boolean isTriangle(int a, int b, int c) {
        /**
         * TODO: You need to complete this method to determine whether  a
         * triangle is formed or not when given the input edge a, b and c.
         */
        if (a <= 0 || b <= 0 || c <= 0) {
            return false;
        }
        if ((a + b > c) && (b + c > a) && (a + c > b)) {
            return true;
        } else {
            return false;
        }
    }

    public boolean isBirthday(int year, int month, int day) {
        /**
         * TODO: You need to complete this method to determine whether a
         * legitimate date of birth between 1990/01/01 and 2019/10/01 is
         * formed or not when given the input year, month and day.
         */
        if (year < 1990 || year > 2019) {
            return false;
        }
        if (year == 2019) {
            if (month > 10) {
                return false;
            } else if (month == 10) {
                if (day == 1) {
                    return true;
                } else {
                    return false;
                }
            }
        }
        List<Integer> list = new ArrayList<Integer>();
        for (int i = 1992; i < 2019; i += 4) {
            list.add(i);
        }
        if (list.contains(year)) {
            switch (month) {
                case 1:
                case 3:
                case 5:
                case 8:
                case 7:
                case 10:
                case 12:
                    if (day >= 1 && day <= 31) {
                        return true;
                    } else {
                        return false;
                    }
                case 2:
                    if (day >= 1 && day <= 29) {
                        return true;
                    } else {
                        return false;
                    }
                case 4:
                case 6:
                case 9:
                case 11:
                    if (day >= 1 && day <= 30) {
                        return true;
                    } else {
                        return false;
                    }
                default:
                    return false;
            }
        } else {
            switch (month) {
                case 1:
                case 3:
                case 5:
                case 8:
                case 7:
                case 10:
                case 12:
                    if (day >= 1 && day <= 31) {
                        return true;
                    } else {
                        return false;
                    }
                case 2:
                    if (day >= 1 && day <= 28) {
                        return true;
                    } else {
                        return false;
                    }
                case 4:
                case 6:
                case 9:
                case 11:
                    if (day >= 1 && day <= 30) {
                        return true;
                    } else {
                        return false;
                    }
                default:
                    return false;
            }

        }

    }

    public Double miniCalculator(double a, double b, char op) {
        /**
         * TODO: You need to complete this method to form a small calculator which
         * can calculate the formula: "a op b", the op here can be four basic
         * operation: "+","-","*","/".
         */
        if (op == '+') {
            return a + b;
        } else if (op == '-') {
            return a - b;
        } else if (op == '*') {
            return a * b;
        } else if (op == '/') {
            return a / b;
        } else {
            throw new ArithmeticException("error op");
        }
    }

}
