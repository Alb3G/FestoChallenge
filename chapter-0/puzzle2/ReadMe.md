Example {
	X; X; R; Q; -> 0
	X; Y; R; Q; -> 1
	Y; X; R; Q; -> 1
	Y; Y; R; Q; -> 1

	X; X; N; Q; -> 0
	Y; X; N; Q; -> 0
	X; Y; N; Q; -> 0
	Y; Y; N; Q; -> 1
}

Problem to Solve {
	YXYY; YYXY; N; Q;
	XYXX; XYYX; R; Q;
}

I didnt need any code to solve this one, I solved as a truth table, this was my approach to solve the problem:

- First line (YXYY; YYXY; N; Q;) its an And(&&) gate so we only get 1 when the 2 entries before N and Q are YY,
so If we check each entry like this (Y)XYY; (Y)YXY;, the output will be 1, next step Y(X)YY; Y(Y)XY;, here the 
output would be 0, and repeat this process with each entry on both lines and then concatenate the result.

so the result to this puzzle is "10010110"