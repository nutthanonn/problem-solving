def isValid(s):
    OpenSymbol = ["(", "[", "{"]
    queue = []
    checkQueue = {
        ")": "(", 
        "]": "[", 
        "}": "{"
    }
    closeSymbol = [")", "]", "}"]
    if s[0] in closeSymbol or len(s)%2 == 1:
        return False 

    for i in s:
        if i in closeSymbol and len(queue) == 0:
            return False

        if i in OpenSymbol:
            queue.append(i)
        elif i in closeSymbol and checkQueue[i] == queue[-1]:
            queue.pop(len(queue)-1)
        else:
            queue.append(i)

    if len(queue) > 0:
        return False
    else:
        return True
        
print(isValid("(){}}{"))