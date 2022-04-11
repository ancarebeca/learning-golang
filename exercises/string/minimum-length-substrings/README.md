Minimum Length Substrings

    You are given two strings s and t. You can select any substring of string s and rearrange the characters of the 
    selected substring. Determine the minimum length of the substring of s such that string t is a substring of the selected substring.

    Signature
        int minLengthSubstring(String s, String t)
    
    Input
        s and t are non-empty strings that contain less than 1,000,000 characters each
    
    Output
        Return the minimum length of the substring of s. If it is not possible, return -1
    
    Example
        s = "dcbefebce"
        t = "fd"
        output = 5
    
    Explanation:
        Substring "dcbef" can be rearranged to "cfdeb", "cefdb", and so on. String t is a substring of "cfdeb". Thus, the minimum 
        length required is 5.

    Do you need some help?
        - https://go.dev/blog/maps
        - https://labuladong.gitbook.io/algo-en/iii.-algorithmic-thinking/slidingwindowtechnique
        - https://www.youtube.com/watch?v=jSto0O4AJbM   
        - https://itnext.io/sliding-window-algorithm-technique-6001d5fbe8b3
        - https://www.baeldung.com/cs/sliding-window-algorithm