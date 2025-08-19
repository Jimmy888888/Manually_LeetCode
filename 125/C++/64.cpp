#include <cctype>
#include <algorithm>

using namespace std;

class Solution {
public:
    bool isPalindrome(string s) {
        string filter_str = "";

        // turn into lower case and remove other non-alphanumeric characters
        for (char c: s){
            if(isalpha(static_cast<unsigned char>(c))){
                filter_str += tolower(static_cast<unsigned char>(c));
            }
            else if(isdigit(static_cast<unsigned char>(c))){
                filter_str += c;
            }
        }

        // handle "" case
        if (filter_str.length() == 0){
            return true;
        }

        int mind_len = filter_str.length()/2;
        int mind_potion = mind_len;
        
        if (filter_str.length()%2 != 0){
            mind_potion = mind_len + 1;
        }

        string front_str = filter_str.substr(0, mind_len);
        string end_str = filter_str.substr(mind_potion);
        reverse(end_str.begin(), end_str.end());

        // cout << front_str << "\n";
        // cout << end_str << "\n";

        if (front_str != end_str){
            return false;
        }

        return true;
    }
};