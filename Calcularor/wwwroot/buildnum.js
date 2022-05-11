function build(num) {
    buildNum(num);
}

function buildNum(num) {
    if (num == "(") {
        s.push(num);
    }
    if (num == ")" && checkSymbol()) {
        document.getElementById("screen").innerHTML += num;
    } else if (num != ")") {
        document.getElementById("screen").innerHTML += num;
        document.getElementById("screen").innerHTML = checkMul(document.getElementById("screen").innerHTML);
        document.getElementById("screen").innerHTML = checkOp(document.getElementById("screen").innerHTML);
    }

    if (document.getElementById("screen").innerHTML.length > 22) {
        document.getElementById("screen").style.fontSize = 'medium';
    } else {
        document.getElementById("screen").style.fontSize = '32px';
    }
}

function buildBracket() {
    document.getElementById("screen").innerHTML += "()";
}

function del() {
    var value = document.getElementById('screen').innerHTML;
    if (value.charAt(value.length - 1) == "(") {
        s.pop();
    } else if (value.charAt(value.length - 1) == ")") {
        s.push("(");
    }
    document.getElementById('screen').innerHTML = value.substr(0, value.length - 1);
}

function reset() {
    document.getElementById('screen').innerHTML = "";
    this.s = [];
}

function isOperand(str) {
    if (str == "*" || str == "-" || str == "+" || str == "/" || str == "^" || str == "(" || str == ")" || str == "√" || str == ".") {
        return false;
    } else {
        return true;
    }
}

var s = [];

function checkMul(str) {
    var prev = "";
    var sub = "";
    for (i = 0; i < str.length; i++) {
        if (str.charAt(i) == "(" && i != 0 && isOperand(prev)) {
            sub = str.substr(i, str.length);
            str = str.substr(0, i) + "*" + sub;
        } else if (prev == ")" && isOperand(str.charAt(i))) {
            sub = str.substr(i, str.length);
            str = str.substr(0, i) + "*" + sub;
        } else if (prev == ")" && str.charAt(i) == "(") {
            sub = str.substr(i, str.length);
            str = str.substr(0, i) + "*" + sub;
        }
        prev = str.charAt(i);
    }
    return str;
}

function checkSymbol(str) {
    if (s.pop() == "(") {
        return true;
    }
}

function checkOp(str) {
    var prev = str.charAt(str.length - 2);
    var pre = str.charAt(str.length - 1);
    if (pre != "(" && pre != ")" && prev != ")") {
        if (!isOperand(prev) && prev != "√" && !isOperand(pre)) {
            str = str.substr(0, str.length - 2) + pre;
        }
    }
    return str;
}