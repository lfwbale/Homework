package com.test.jacoco.test;

import static org.junit.Assert.assertEquals;
import static org.junit.Assert.assertThat;

/**
 * Created by superZh on 2019/10/1.
 */

import org.junit.Rule;
import org.junit.Test;
import com.test.jacoco.HelloWorld;
import org.junit.rules.ExpectedException;

public class HelloWorldTest {

    @Rule
    public ExpectedException thrown = ExpectedException.none();

    @Test
    public void testMethod1() {
        HelloWorld hw = new HelloWorld();
        String a = hw.Method1();
        assertEquals("Hello World", a);
    }

    @Test
    public void testMethod2() {
        HelloWorld hw = new HelloWorld();
        int a = hw.Method2(2, 1);
        assertEquals(3, a);
    }

    @Test
    public void testMethod3() {
        /**
         * TODO：finish the test function
         */
        HelloWorld hw = new HelloWorld();
        // 第一个条件分支
        int a = hw.Method3(6, -1, 1); // true true
        int b = hw.Method3(6, -1, 1);  // true false = true
        int c = hw.Method3(6, 1, 1); // false true = true
        int d = hw.Method3(4, -1, 1); // false true = true
        int e = hw.Method3(4, 1, 1); // false true = true
        int f = hw.Method3(6, 1, -1); // false false = false
        int g = hw.Method3(4, -1, -1); // false false = false
        int h = hw.Method3(4, 1, -1); // false false = false
        // 第二个条件分支
        int i = hw.Method3(6, 1, -3); // false false = false
        int j = hw.Method3(4, -1, -3); // true true = true
        int k = hw.Method3(4, 1, -1); // true false = false


    }

    /**
     * TODO: add the test function of other methods in HelloWorld.java
     */
    @Test
    public void testMethod4() {
        HelloWorld hw = new HelloWorld();
        int a = hw.Method4(0, 1, 1, 1, 1);
        int b = hw.Method4(1, 1, 0, 1, 1);
        int c = hw.Method4(5, 1, 1, 1, 0);
        int d = hw.Method4(5, 5, 1, 1, 0);
        try{
            int e = hw.Method4(1, 2, 1, 1, 0);
        }catch (Exception err){
            System.out.println("error");
        }
        try{
            int f = hw.Method4(5, 1, 1, 2, 0);;
        }catch (Exception err){
            System.out.println("error");
        }
        try{
            int g = hw.Method4(1, 1, 2, 2, 0);
        }catch (Exception err){
            System.out.println("error");
        }
        try{
            int h = hw.Method4(1, 2, 1, 2, 0);
        }catch (Exception err){
            System.out.println("error");
        }
    }

    @Test
    public void testBug() {
        HelloWorld hw = new HelloWorld();
        boolean f1 = hw.bug(5);
        assertEquals(true, f1);
        boolean f2 = hw.bug(6);
        assertEquals(false, f2);
    }

    @Test
    public void testisTriangle() {
        HelloWorld hw = new HelloWorld();
        boolean f0 = hw.isTriangle(3, 4, 5);
        assertEquals(f0, true);
        boolean f1 = hw.isTriangle(0, 1, 1);
        assertEquals(f1, false);
        boolean f2 = hw.isTriangle(1, 0, 1);
        assertEquals(f2, false);
        boolean f3 = hw.isTriangle(1, 1, 0);
        assertEquals(f3, false);
        boolean f4 = hw.isTriangle(1, 1, 2);
        assertEquals(f4, false);
        boolean f5 = hw.isTriangle(4, 2, 2);
        assertEquals(f5, false);
        boolean f6 = hw.isTriangle(2, 5, 2);
        assertEquals(f6, false);

    }

    @Test
    public void testBirthday() {
        HelloWorld hw = new HelloWorld();
        assertEquals(hw.isBirthday(1989, 1, 1), false);
        assertEquals(hw.isBirthday(2020, 1, 1), false);
        assertEquals(hw.isBirthday(2019, 11, 1), false);
        assertEquals(hw.isBirthday(2019, 10, 1), true);
        assertEquals(hw.isBirthday(2019, 10, 2), false);
        assertEquals(hw.isBirthday(2019, 9, 1), true);
        assertEquals(hw.isBirthday(1992, 1, 1), true);
        assertEquals(hw.isBirthday(1992, 1, -1), false);
        assertEquals(hw.isBirthday(1992, 1, 32), false);
        assertEquals(hw.isBirthday(1992, 3, 1), true);
        assertEquals(hw.isBirthday(1992, 5, 1), true);
        assertEquals(hw.isBirthday(1992, 7, 1), true);
        assertEquals(hw.isBirthday(1992, 8, 1), true);
        assertEquals(hw.isBirthday(1992, 10, 1), true);
        assertEquals(hw.isBirthday(1992, 12, 1), true);
        assertEquals(hw.isBirthday(1992, 2, -1), false);
        assertEquals(hw.isBirthday(1992, 2, 1), true);
        assertEquals(hw.isBirthday(1992, 2, 29), true);
        assertEquals(hw.isBirthday(1992, 2, 30), false);
        assertEquals(hw.isBirthday(1992, 4, 1), true);
        assertEquals(hw.isBirthday(1992, 4, -1), false);
        assertEquals(hw.isBirthday(1992, 4, 31), false);
        assertEquals(hw.isBirthday(1992, 6, 1), true);
        assertEquals(hw.isBirthday(1992, 9, 1), true);
        assertEquals(hw.isBirthday(1992, 11, 1), true);
        assertEquals(hw.isBirthday(1992, 13, 1), false);

        assertEquals(hw.isBirthday(1993, 1, 1), true);
        assertEquals(hw.isBirthday(1993, 1, -1), false);
        assertEquals(hw.isBirthday(1993, 1, 32), false);
        assertEquals(hw.isBirthday(1993, 3, 1), true);
        assertEquals(hw.isBirthday(1993, 5, 1), true);
        assertEquals(hw.isBirthday(1993, 7, 1), true);
        assertEquals(hw.isBirthday(1993, 8, 1), true);
        assertEquals(hw.isBirthday(1993, 10, 1), true);
        assertEquals(hw.isBirthday(1993, 12, 1), true);
        assertEquals(hw.isBirthday(1993, 2, -1), false);
        assertEquals(hw.isBirthday(1993, 2, 1), true);
        assertEquals(hw.isBirthday(1993, 2, 29), false);
        assertEquals(hw.isBirthday(1993, 2, 30), false);
        assertEquals(hw.isBirthday(1993, 4, 1), true);
        assertEquals(hw.isBirthday(1993, 4, -1), false);
        assertEquals(hw.isBirthday(1993, 4, 31), false);
        assertEquals(hw.isBirthday(1993, 6, 1), true);
        assertEquals(hw.isBirthday(1993, 9, 1), true);
        assertEquals(hw.isBirthday(1993, 11, 1), true);
        assertEquals(hw.isBirthday(1993, 13, 1), false);
    }

    @Test
    public void testCalculate() {
        HelloWorld hw = new HelloWorld();
        double a = hw.miniCalculator(1, 1, '+');
        double b = hw.miniCalculator(1, 1, '-');
        double c = hw.miniCalculator(1, 1, '*');
        double d = hw.miniCalculator(1, 1, '/');
        thrown.expect(ArithmeticException.class);
        double e = hw.miniCalculator(1, 1, '%');
    }
}
