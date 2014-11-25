package jsbeautifier

import (
	"jsbeautifier/optargs"
	"testing"
)

// Copyright (c) 2014 Ditashi Sayomi

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

var ts *testing.T

var test_options optargs.MapType

func TestBeautifier(t *testing.T) {
	ts = t

	test_options = default_options

	test_options["indent_size"] = 4
	test_options["indent_char"] = " "
	test_options["preserve_newlines"] = true
	test_options["jslint_happy"] = false
	test_options["keep_array_indentation"] = false
	test_options["brace_style"] = "collapse"

	// Unicode Support
	test("var \u0CA0_\u0CA0 = \"hi\";")
	test("var \u00e4 x = {\n    \u00e4rgerlich: true\n};")

	// End With Newline - (eof = "\n")
	test_options["end_with_newline"] = true
	test("", "\n")
	test("   return .5", "   return .5\n")
	test("   \n\nreturn .5\n\n\n\n", "   return .5\n")
	test("\n")

	// End With Newline - (eof = "")
	test_options["end_with_newline"] = false
	test("")
	test("   return .5")
	test("   \n\nreturn .5\n\n\n\n", "   return .5")
	test("\n", "")

	// New Test Suite

	// Old tests
	test("")
	test("   return .5")
	test("   return .5;\n   a();")
	test("    return .5;\n    a();")
	test("     return .5;\n     a();")
	test("   < div")
	test("a        =          1", "a = 1")
	test("a=1", "a = 1")
	test("(3) / 2")
	test("[\"a\", \"b\"].join(\"\")")
	test("a();\n\nb();")
	test("var a = 1 var b = 2", "var a = 1\nvar b = 2")
	test("var a=1, b=c[d], e=6;", "var a = 1,\n    b = c[d],\n    e = 6;")
	test("var a,\n    b,\n    c;")
	test("let a = 1 let b = 2", "let a = 1\nlet b = 2")
	test("let a=1, b=c[d], e=6;", "let a = 1,\n    b = c[d],\n    e = 6;")
	test("let a,\n    b,\n    c;")
	test("const a = 1 const b = 2", "const a = 1\nconst b = 2")
	test("const a=1, b=c[d], e=6;", "const a = 1,\n    b = c[d],\n    e = 6;")
	test("const a,\n    b,\n    c;")
	test("a = \" 12345 \"")
	test("a = ' 12345 '")
	test("if (a == 1) b = 2;")
	test("if(1){2}else{3}", "if (1) {\n    2\n} else {\n    3\n}")
	test("if(1||2);", "if (1 || 2);")
	test("(a==1)||(b==2)", "(a == 1) || (b == 2)")
	test("var a = 1 if (2) 3;", "var a = 1\nif (2) 3;")
	test("a = a + 1")
	test("a = a == 1")
	test("/12345[^678]*9+/.match(a)")
	test("a /= 5")
	test("a = 0.5 * 3")
	test("a *= 10.55")
	test("a < .5")
	test("a <= .5")
	test("a<.5", "a < .5")
	test("a<=.5", "a <= .5")
	test("a = 0xff;")
	test("a=0xff+4", "a = 0xff + 4")
	test("a = [1, 2, 3, 4]")
	test("F*(g/=f)*g+b", "F * (g /= f) * g + b")
	test("a.b({c:d})", "a.b({\n    c: d\n})")
	test("a.b\n(\n{\nc:\nd\n}\n)", "a.b({\n    c: d\n})")
	test("a.b({c:\"d\"})", "a.b({\n    c: \"d\"\n})")
	test("a.b\n(\n{\nc:\n\"d\"\n}\n)", "a.b({\n    c: \"d\"\n})")
	test("a=!b", "a = !b")
	test("a=!!b", "a = !!b")
	test("a?b:c", "a ? b : c")
	test("a?1:2", "a ? 1 : 2")
	test("a?(b):c", "a ? (b) : c")
	test("x={a:1,b:w==\"foo\"?x:y,c:z}", "x = {\n    a: 1,\n    b: w == \"foo\" ? x : y,\n    c: z\n}")
	test("x=a?b?c?d:e:f:g;", "x = a ? b ? c ? d : e : f : g;")
	test("x=a?b?c?d:{e1:1,e2:2}:f:g;", "x = a ? b ? c ? d : {\n    e1: 1,\n    e2: 2\n} : f : g;")
	test("function void(void) {}")
	test("if(!a)foo();", "if (!a) foo();")
	test("a=~a", "a = ~a")
	test("a;/*comment*/b;", "a; /*comment*/\nb;")
	test("a;/* comment */b;", "a; /* comment */\nb;")

	// simple comments don't get touched at all
	test("a;/*\ncomment\n*/b;", "a;\n/*\ncomment\n*/\nb;")
	test("a;/**\n* javadoc\n*/b;", "a;\n/**\n * javadoc\n */\nb;")
	test("a;/**\n\nno javadoc\n*/b;", "a;\n/**\n\nno javadoc\n*/\nb;")

	// comment blocks detected and reindented even w/o javadoc starter
	test("a;/*\n* javadoc\n*/b;", "a;\n/*\n * javadoc\n */\nb;")
	test("if(a)break;", "if (a) break;")
	test("if(a){break}", "if (a) {\n    break\n}")
	test("if((a))foo();", "if ((a)) foo();")
	test("for(var i=0;;) a", "for (var i = 0;;) a")
	test("for(var i=0;;)\na", "for (var i = 0;;)\n    a")
	test("a++;")
	test("for(;;i++)a()", "for (;; i++) a()")
	test("for(;;i++)\na()", "for (;; i++)\n    a()")
	test("for(;;++i)a", "for (;; ++i) a")
	test("return(1)", "return (1)")
	test("try{a();}catch(b){c();}finally{d();}", "try {\n    a();\n} catch (b) {\n    c();\n} finally {\n    d();\n}")

	//  magic function call
	test("(xx)()")

	// another magic function call
	test("a[1]()")
	test("if(a){b();}else if(c) foo();", "if (a) {\n    b();\n} else if (c) foo();")
	test("switch(x) {case 0: case 1: a(); break; default: break}", "switch (x) {\n    case 0:\n    case 1:\n        a();\n        break;\n    default:\n        break\n}")
	test("switch(x){case -1:break;case !y:break;}", "switch (x) {\n    case -1:\n        break;\n    case !y:\n        break;\n}")
	test("a !== b")
	test("if (a) b(); else c();", "if (a) b();\nelse c();")

	// typical greasemonkey start
	test("// comment\n(function something() {})")

	// duplicating newlines
	test("{\n\n    x();\n\n}")
	test("if (a in b) foo();")
	test("if(X)if(Y)a();else b();else c();", "if (X)\n    if (Y) a();\n    else b();\nelse c();")
	test("if (foo) bar();\nelse break")
	test("var a, b;")
	test("var a = new function();")
	test("new function")
	test("var a, b")
	test("{a:1, b:2}", "{\n    a: 1,\n    b: 2\n}")
	test("a={1:[-1],2:[+1]}", "a = {\n    1: [-1],\n    2: [+1]\n}")
	test("var l = {'a':'1', 'b':'2'}", "var l = {\n    'a': '1',\n    'b': '2'\n}")
	test("if (template.user[n] in bk) foo();")
	test("return 45")
	test("return this.prevObject ||\n\n    this.constructor(null);")
	test("If[1]")
	test("Then[1]")
	test("a = 1e10")
	test("a = 1.3e10")
	test("a = 1.3e-10")
	test("a = -1.3e-10")
	test("a = 1e-10")
	test("a = e - 10")
	test("a = 11-10", "a = 11 - 10")
	test("a = 1;// comment", "a = 1; // comment")
	test("a = 1; // comment")
	test("a = 1;\n // comment", "a = 1;\n// comment")
	test("a = [-1, -1, -1]")

	// The exact formatting these should have is open for discussion, but they are at least reasonable
	test("a = [ // comment\n    -1, -1, -1\n]")
	test("var a = [ // comment\n    -1, -1, -1\n]")
	test("a = [ // comment\n    -1, // comment\n    -1, -1\n]")
	test("var a = [ // comment\n    -1, // comment\n    -1, -1\n]")
	test("o = [{a:b},{c:d}]", "o = [{\n    a: b\n}, {\n    c: d\n}]")

	// was: extra space appended
	test("if (a) {\n    do();\n}")

	// if/else statement with empty body
	test("if (a) {\n// comment\n}else{\n// comment\n}", "if (a) {\n    // comment\n} else {\n    // comment\n}")

	// multiple comments indentation
	test("if (a) {\n// comment\n// comment\n}", "if (a) {\n    // comment\n    // comment\n}")
	test("if (a) b() else c();", "if (a) b()\nelse c();")
	test("if (a) b() else if c() d();", "if (a) b()\nelse if c() d();")
	test("{}")
	test("{\n\n}")
	test("do { a(); } while ( 1 );", "do {\n    a();\n} while (1);")
	test("do {} while (1);")
	test("do {\n} while (1);", "do {} while (1);")
	test("do {\n\n} while (1);")
	test("var a = x(a, b, c)")
	test("delete x if (a) b();", "delete x\nif (a) b();")
	test("delete x[x] if (a) b();", "delete x[x]\nif (a) b();")
	test("for(var a=1,b=2)d", "for (var a = 1, b = 2) d")
	test("for(var a=1,b=2,c=3) d", "for (var a = 1, b = 2, c = 3) d")
	test("for(var a=1,b=2,c=3;d<3;d++)\ne", "for (var a = 1, b = 2, c = 3; d < 3; d++)\n    e")
	test("function x(){(a||b).c()}", "function x() {\n    (a || b).c()\n}")
	test("function x(){return - 1}", "function x() {\n    return -1\n}")
	test("function x(){return ! a}", "function x() {\n    return !a\n}")
	test("x => x")
	test("(x) => x")
	test("x => { x }", "x => {\n    x\n}")
	test("(x) => { x }", "(x) => {\n    x\n}")

	// a common snippet in jQuery plugins
	test("settings = $.extend({},defaults,settings);", "settings = $.extend({}, defaults, settings);")
	test("$http().then().finally().default()")
	test("$http()\n.then()\n.finally()\n.default()", "$http()\n    .then()\n    .finally()\n    .default()")
	test("$http().when.in.new.catch().throw()")
	test("$http()\n.when\n.in\n.new\n.catch()\n.throw()", "$http()\n    .when\n    .in\n    .new\n    .catch()\n    .throw()")
	test("{xxx;}()", "{\n    xxx;\n}()")
	test("a = 'a'\nb = 'b'")
	test("a = /reg/exp")
	test("a = /reg/")
	test("/abc/.test()")
	test("/abc/i.test()")
	test("{/abc/i.test()}", "{\n    /abc/i.test()\n}")
	test("var x=(a)/a;", "var x = (a) / a;")
	test("x != -1")
	test("for (; s-->0;)t", "for (; s-- > 0;) t")
	test("for (; s++>0;)u", "for (; s++ > 0;) u")
	test("a = s++>s--;", "a = s++ > s--;")
	test("a = s++>--s;", "a = s++ > --s;")
	test("{x=#1=[]}", "{\n    x = #1=[]\n}")
	test("{a:#1={}}", "{\n    a: #1={}\n}")
	test("{a:#1#}", "{\n    a: #1#\n}")
	test("\"incomplete-string")
	test("'incomplete-string")
	test("/incomplete-regex")
	test("`incomplete-template-string")
	test("{a:1},{a:2}", "{\n    a: 1\n}, {\n    a: 2\n}")
	test("var ary=[{a:1}, {a:2}];", "var ary = [{\n    a: 1\n}, {\n    a: 2\n}];")
	// incomplete
	test("{a:#1", "{\n    a: #1")

	// incomplete
	test("{a:#", "{\n    a: #")

	// incomplete
	test("}}}", "}\n}\n}")
	test("<!--\nvoid();\n// -->")

	// incomplete regexp
	/*test("a=/regexp", "a = /regexp")
	test("{a:#1=[],b:#1#,c:#999999#}", "{\n    a: #1=[],\n    b: #1#,\n    c: #999999#\n}")
	test("a = 1e+2")
	test("a = 1e-2")
	test("do{x()}while(a>1)", "do {\n    x()\n} while (a > 1)")
	test("x(); /reg/exp.match(something)", "x();\n/reg/exp.match(something)")
	test("something();(", "something();\n(")
	test("#!she/bangs, she bangs\nf=1", "#!she/bangs, she bangs\n\nf = 1")
	test("#!she/bangs, she bangs\n\nf=1", "#!she/bangs, she bangs\n\nf = 1")*/
	//test("#!she/bangs, she bangs\n\n/* comment */")
	//test("#!she/bangs, she bangs\n\n\n/* comment */")
	/*test("#")
	test("#!")
	test("function namespace::something()")
	test("<!--\nsomething();\n-->")
	test("<!--\nif(i<0){bla();}\n-->", "<!--\nif (i < 0) {\n    bla();\n}\n-->")
	test("{foo();--bar;}", "{\n    foo();\n    --bar;\n}")
	test("{foo();++bar;}", "{\n    foo();\n    ++bar;\n}")
	test("{--bar;}", "{\n    --bar;\n}")
	test("{++bar;}", "{\n    ++bar;\n}")
	test("if(true)++a;", "if (true) ++a;")
	test("if(true)\n++a;", "if (true)\n    ++a;")
	test("if(true)--a;", "if (true) --a;")
	test("if(true)\n--a;", "if (true)\n    --a;")*/

	// Handling of newlines around unary ++ and -- operators
	test("{foo\n++bar;}", "{\n    foo\n    ++bar;\n}")
	test("{foo++\nbar;}", "{\n    foo++\n    bar;\n}")

	// This is invalid, but harder to guard against. Issue #203.
	test("{foo\n++\nbar;}", "{\n    foo\n    ++\n    bar;\n}")

	// regexps
	test("a(/abc\\/\\/def/);b()", "a(/abc\\/\\/def/);\nb()")
	test("a(/a[b\\[\\]c]d/);b()", "a(/a[b\\[\\]c]d/);\nb()")

	// incomplete char class
	test("a(/a[b\\[")

	// allow unescaped / in char classes
	test("a(/[a/b]/);b()", "a(/[a/b]/);\nb()")
	test("typeof /foo\\//;")
	test("yield /foo\\//;")
	test("throw /foo\\//;")
	test("do /foo\\//;")
	test("return /foo\\//;")
	test("switch (a) {\n    case /foo\\//:\n        b\n}")
	test("if (a) /foo\\//\nelse /foo\\//;")
	test("if (foo) /regex/.test();")
	test("function foo() {\n    return [\n        \"one\",\n        \"two\"\n    ];\n}")
	test("a=[[1,2],[4,5],[7,8]]", "a = [\n    [1, 2],\n    [4, 5],\n    [7, 8]\n]")
	test("a=[[1,2],[4,5],function(){},[7,8]]", "a = [\n    [1, 2],\n    [4, 5],\n    function() {},\n    [7, 8]\n]")
	test("a=[[1,2],[4,5],function(){},function(){},[7,8]]", "a = [\n    [1, 2],\n    [4, 5],\n    function() {},\n    function() {},\n    [7, 8]\n]")
	test("a=[[1,2],[4,5],function(){},[7,8]]", "a = [\n    [1, 2],\n    [4, 5],\n    function() {},\n    [7, 8]\n]")
	test("a=[b,c,function(){},function(){},d]", "a = [b, c, function() {}, function() {}, d]")
	test("a=[b,c,\nfunction(){},function(){},d]", "a = [b, c,\n    function() {},\n    function() {},\n    d\n]")
	test("a=[a[1],b[4],c[d[7]]]", "a = [a[1], b[4], c[d[7]]]")
	test("[1,2,[3,4,[5,6],7],8]", "[1, 2, [3, 4, [5, 6], 7], 8]")

	test("[[[\"1\",\"2\"],[\"3\",\"4\"]],[[\"5\",\"6\",\"7\"],[\"8\",\"9\",\"0\"]],[[\"1\",\"2\",\"3\"],[\"4\",\"5\",\"6\",\"7\"],[\"8\",\"9\",\"0\"]]]", "[\n    [\n        [\"1\", \"2\"],\n        [\"3\", \"4\"]\n    ],\n    [\n        [\"5\", \"6\", \"7\"],\n        [\"8\", \"9\", \"0\"]\n    ],\n    [\n        [\"1\", \"2\", \"3\"],\n        [\"4\", \"5\", \"6\", \"7\"],\n        [\"8\", \"9\", \"0\"]\n    ]\n]")
	test("{[x()[0]];indent;}", "{\n    [x()[0]];\n    indent;\n}")

	test("{{}/z/}", "{\n    {}\n    /z/\n}")
	test("return ++i", "return ++i")
	test("return !!x", "return !!x")
	test("return !x", "return !x")
	test("return [1,2]", "return [1, 2]")
	test("return;", "return;")
	test("return\nfunc", "return\nfunc")
	test("catch(e)", "catch (e)")
	test("yield [1, 2]")

	test("var a=1,b={foo:2,bar:3},{baz:4,wham:5},c=4;",
		"var a = 1,\n    b = {\n        foo: 2,\n        bar: 3\n    },\n    {\n        baz: 4,\n        wham: 5\n    }, c = 4;")
	test("var a=1,b={foo:2,bar:3},{baz:4,wham:5},\nc=4;",
		"var a = 1,\n    b = {\n        foo: 2,\n        bar: 3\n    },\n    {\n        baz: 4,\n        wham: 5\n    },\n    c = 4;")

	// inline comment
	test("function x(/*int*/ start, /*string*/ foo)", "function x( /*int*/ start, /*string*/ foo)")

	// javadoc comment
	test("/**\n* foo\n*/", "/**\n * foo\n */")
	test("{\n/**\n* foo\n*/\n}", "{\n    /**\n     * foo\n     */\n}")

	// starless block comment
	test("/**\nfoo\n*/")
	test("/**\nfoo\n**/")
	test("/**\nfoo\nbar\n**/")
	test("/**\nfoo\n\nbar\n**/")
	test("/**\nfoo\n    bar\n**/")
	test("{\n/**\nfoo\n*/\n}", "{\n    /**\n    foo\n    */\n}")
	test("{\n/**\nfoo\n**/\n}", "{\n    /**\n    foo\n    **/\n}")
	test("{\n/**\nfoo\nbar\n**/\n}", "{\n    /**\n    foo\n    bar\n    **/\n}")
	test("{\n/**\nfoo\n\nbar\n**/\n}", "{\n    /**\n    foo\n\n    bar\n    **/\n}")
	test("{\n/**\nfoo\n    bar\n**/\n}", "{\n    /**\n    foo\n        bar\n    **/\n}")
	test("{\n    /**\n    foo\nbar\n    **/\n}")

	test("var a,b,c=1,d,e,f=2;", "var a, b, c = 1,\n    d, e, f = 2;")
	test("var a,b,c=[],d,e,f=2;", "var a, b, c = [],\n    d, e, f = 2;")
	test("function() {\n    var a, b, c, d, e = [],\n        f;\n}")

	test("do/regexp/;\nwhile(1);", "do /regexp/;\nwhile (1);")

	test("var a = a,\na;\nb = {\nb\n}", "var a = a,\n    a;\nb = {\n    b\n}")

	test("var a = a,\n    /* c */\n    b;")
	test("var a = a,\n    // c\n    b;")

	test("foo.(\"bar\");")

	test("if (a) a()\nelse b()\nnewline()")
	test("if (a) a()\nnewline()")
	test("a=typeof(x)", "a = typeof(x)")

	test("var a = function() {\n        return null;\n    },\n    b = false;")

	test("var a = function() {\n    func1()\n}")
	test("var a = function() {\n    func1()\n}\nvar b = function() {\n    func2()\n}")

	// Code with and without semicolons

}

func test(options ...string) {
	if len(options) == 1 {
		assertMatch(options[0], options[0])
	} else if len(options) == 2 {
		assertMatch(options[0], options[1])
	} else {
		ts.Error("Cannot test for nothing or more than 1 input")
	}
}

func assertMatch(input, expect string) {

	result, _ := Beautify(&input, test_options)

	if result != expect {
		ts.Error("Input", input, "Result: ", result, " did not match ", expect)
	}

}
