#include <iostream>
#include <iomanip>
#include <limits>
#include<string>
using namespace std;

int main() {
    int i = 4;
    double d = 4.0;
    string s = "HackerRank ";

    
    // Declare second integer, double, and String variables.
    int my_int;
    double my_double;
    string my_string;
    // Read and save an integer, double, and String to your variables.
    
    
    
    
    // Print the sum of both integer variables on a new line.
    cin >> my_int ;
    cout<< my_int + i<< "\n";
    // Print the sum of the double variables on a new line.
    cin >> my_double;
    cout << my_double + d<< "\n";
    // Concatenate and print the String variables on a new line
    getline(cin, my_string);
    cout << s + my_string << "\n";
    // The 's' variable above should be printed first.

    return 0;
}