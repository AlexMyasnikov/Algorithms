Console.WriteLine("Hello, World!");

char[] open = {'{', '[', '('};
char[] close = {'}', ']', ')'};

string[] lines = System.IO.File.ReadAllLines("exp.txt");
foreach (string l in lines) {
    Stack<char> st = new Stack<char>();
    foreach (char el in l.ToCharArray()) {
        if (open.Contains(el)) {
            st.Push(el);
        } else if (close.Contains(el)) {
            int idx_c = Array.IndexOf(close, el);
            int idx_o = Array.IndexOf(open, st.Peek());
            if (idx_c == idx_o) {
                st.Pop();
            } else {
                break;
            }
        }
    }

    if (st.Count() != 0) {
        Console.WriteLine("{0} – неправильная скобочная последовательность", l);
    } else {
        Console.WriteLine("{0} – правильная скобочная последовательность", l);
    }
}

// ({})