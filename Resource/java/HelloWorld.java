public class HelloWorld {
    public static void main(String []args) {
        System.out.println("Hello CoderHG");


        System.out.println(reverse("Hello CoderHG 123456"));
    }

    public static String reverse(String str) {
        String[] chars = str.split(" ");
        for (int i = 0; i < chars.length; i++) {
            if (chars[i] != null) {
                chars[i] = reverse(chars[i].toCharArray());
            }
        }

        for (int i = 0; i < chars.length; i++) {
        	System.out.println(chars[i]);
        }
        StringBuffer s4 = new StringBuffer();
        for (String string: chars) {
        	s4.append(string + " ");
        }

        return s4.toString();
        // return StringUtils.join(chars);
        // return Arrays.toString(chars);
        // return "new String(chars)";
    }

    private static String reverse(char[] arr) {
        int length = arr.length;
        if (length == 0) {
            return "";
        }

        if (length == 1) {
            return arr[0] + "";
        }

        int start = 0;
        int end = arr.length - 1;
        while (start < end) {
            char temp = arr[start];
            arr[start] = arr[end];
            arr[end] = temp;
            start++;
            end--;
        }

        return new String(arr);
    }
}