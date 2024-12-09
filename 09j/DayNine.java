import java.util.*;
import java.nio.file.*;


class Main {

    // need to have index for part 2
    private static int allocate(List<PriorityQueue<Integer>> free, int size, int cur) {
        int best = -1;
        int bestIndex = cur;
        for (int i = size; i < free.size(); i++) {
            if (free.get(i).isEmpty() || free.get(i).peek() > bestIndex) {
                continue;
            }
            best = i;
            bestIndex = free.get(i).peek();
        }
        if (bestIndex < cur) {
            free.get(best).poll();
            int remain = best - size;
            free.get(remain).offer(bestIndex+size);
        }
        return bestIndex;
    }

    private static void part1(List<Long> fs, List<PriorityQueue<Integer>> free) {
        for (int i = fs.size()-1; i >= 0; i--) {
            if (fs.get(i) == -1) {
                continue;
            }
            int j = allocate(free, 1, i);
            if (j >= i) {
                break;
            }
            fs.set(j, fs.get(i));
            fs.set(i, -2L);
        }
        System.out.println(fs);
    }

    private static void part2(List<Long> fs, List<PriorityQueue<Integer>> free) {
        for (int i = fs.size()-1; i >= 0;) {
            if (fs.get(i) < 0) {
                i--;
                continue;
            }
            int j = i;
            long v = fs.get(i);
            while (j >= 0 && v == fs.get(j)) {
                j--;
            }
            int size = i - j;
            int pos = allocate(free, size, i);
            if (pos < i) {
                while (i > j) {
                    fs.set(i--, -2L);
                    fs.set(pos++, v);
                }
            } else {
                i = j;
            }
        }
    }

    private static long score(List<Long> fs) {
        long hash = 0;
        for (int i = 0; i < fs.size(); i++) {
            if (fs.get(i) < 0) {
                continue;
            }
            hash += i * fs.get(i);
        }
        return hash;
    }
    

    public static void main(String[] args) {
        String filename = "input";
        if (args.length >= 1) {
            filename = args[0];
        }
        String s;
        try {
            s = Files.readString(Paths.get(filename));
        } catch (Exception e) {
            throw new IllegalArgumentException("Unable to open input.", e);
        }
        s = s.trim();
        int[] tokens = new int[s.length()];
        for (int i = 0; i < s.length(); i++) {
            tokens[i] = s.charAt(i) - '0';
        }

        List<Long> fs = new ArrayList<>();
        List<PriorityQueue<Integer>> free = new ArrayList<>();
        while (free.size() < 10) {
            free.add(new PriorityQueue<>());
        }

        int pos = 0;
        for (int i = 0; i < tokens.length; i++) {
            int v;
            if ( (i&1) == 0) {
                v = i/2;
            }
            else {
                v = -1;
                free.get(tokens[i]).offer(pos);
            }
            for (int j = 0; j < tokens[i]; j++) {
                fs.add((long) v);
                pos++;
            }
        }

        //part1(fs, free);
        part2(fs, free);
        System.out.println(score(fs));
    }
}
