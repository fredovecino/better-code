# Better Code
This tool is designed to check your code and warn you about bad practices. It uses regular expressions to find patterns of bad code and outputs a log with a message of what's wrong.

There is a .rule file for every language that contains every pattern to check, since all the rules to check are exposed in files you can create your own custom rules to check patterns that you don't want in your code.

# How to use
Just run the executable through the console specifying the file path
````
$ bc -f {filepath}
````

# Custom rules
Rules are composed of a regular expression, a level in the log and a message. You can create a new rule with the following format:
````
regex= {Your regular expression} level= 0 msg= {Message to show when a match is found}
````
Rules are matched through the file extension, for a file named main.js the program will try to find a js.rule to load the patterns.
