import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.Arrays;

public class Main {
    public static void main(String[] args) throws IOException {
        int count1 = 0;
        int count2 = 0;

        BufferedReader in = new BufferedReader(new FileReader(new File("input.txt")));
        String line;
        while ((line = in.readLine()) != null) {
            int[] levels = Arrays.stream(line.split(" "))
                                 .mapToInt(Integer::parseInt)
                                 .toArray();
            
            if (isSafe(levels, 0, 0)) count1++;
            else if (isDampenerSafe(levels)) count2++;
        }

        System.out.println("Solution 1: " + count1); // 591
        System.out.println("Solution 2: " + (count1 + count2)); // 621
    }

    public static boolean isDampenerSafe(int[] levels) {
        int[] a = new int[levels.length - 1];

        for (int index = 0; index < levels.length; index++) {
            for (int i = 0, j = 0; i < levels.length; i++) { // sorry professors for quadratic time complexity
                if (i != index) a[j++] = levels[i];
            }

            if (isSafe(a, 0, 0)) return true;
        }

        return false;
    }

    public static boolean isSafe(int[] levels, int i, int inc) {
        if (i >= levels.length - 1) return true;

        int curr = levels[i];
        int next = levels[i + 1];
        int diff = Math.abs(curr - next);

        if (diff < 1 || diff > 3 || curr == next) return false;

        if (inc == 0) {
            if (curr > next) return isSafe(levels, i + 1, -1); 
            else return isSafe(levels, i + 1, 1); 
        }

        if (curr > next && inc == 1) return false;
        if (curr < next && inc == -1) return false;

        return isSafe(levels, i + 1, inc);
    } 
}