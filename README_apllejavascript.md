# JAVA 语言了解

代码  
```
#!/usr/bin/osascript

#say "Hello 朱鸿."
#beep 2

set x to 25
set y to 2.5

set z to x * y


set str1 to ""
set str2 to "test"

set result to "Nealyyang"


# display dialog result

set ss to str2 & result

set theLenth to the length of ss

set strToNumber to "15" as number

set numToStr to 14 as string

set exampleList to {1, 2, "345", "zhuhong"}


set addToBegin to {"winter"}
set addToEnd to {"summer", "autumn"}
set currentList to {"spring"}
set modifiedList to addToBegin & currentList & addToEnd


set testList to {"11", "22"}
set item 2 of testList to "33"

testList

set myList to {"a"}
set myString to "b"
--set result to myList & (myString as list)
set result to myList & myString

set myString to "neal's personal website is www.nealyang.cn"
set oldDelimiters to AppleScript's text item delimiters
set AppleScript's text item delimiters to " "
set myList to every text item of myString
set AppleScript's text item delimiters to oldDelimiters
get myList

set stringToBeDisplayed to "Neal is pretty boy"
#set tempVar to display dialog stringToBeDisplayed buttons {"So,so", "Don't know", "yes"}
#set theButtonPressed to button returned of tempVar

#get theButtonPressed & "__ok"

set ageEntered to 2
set myAge to 24
if ageEntered is myAge then
	display dialog "You are as old as I am "
	beep
end if
#say "this sentence is spoken anyway"


try
	set x to 1 / 1
on error the error_message number the error_number
	display dialog "Error: " & the error_number & "." & the error_message buttons {"OK"}
end try


#choose folder

tell application "Finder"
	#open folder "Macintosh HD:Users:zhuhong:Desktop:LIVE"
end tell

on test(lala)
	display dialog lala
end test

test("Hello CoderHG.")

on largest(a, b)
	if a > b then
		return a
	end if
	return b
end largest

set theLargetst to largest(4, 6)

set fileInfo to do shell script "cd ~/Desktop; ls"

tell app "Address Book" to get the name of every person



```