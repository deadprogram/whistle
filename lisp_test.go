package main

import (
    "testing"
)

func TestLisp(t *testing.T) {
    // NOTE: one shared global env for test, meaning order matters here!
    env := GlobalEnv()
    for i, tt := range []struct{
        input string
        want string
    }{
        {
            input: "(begin (define r 10) (* pi (* r r)))",
            want:  "314.1592653589793",
        },
        {
            input: "(if (> (* 11 11) 120) (* 7 6) oops)",
            want:  "42",
        },
        {
            input: "(define circle-area (lambda (r) (* pi (* r r))))",
            want:  "0",
        },
        {
            input: "(circle-area 3)",
            want:  "28.274333882308138",
        },
        {
            input: "(quote quoted)",
            want:  "quoted",
        },
        {
            input: "(if (number? (quote ())) 4 5)",
            want:  "5",
        },
        {
            input: "(car (quote (1 2 3)))",
            want:  "1",
        },
        {
            input: "(cdr (quote (1 2 3)))",
            want:  "[2 3]",
        },
        {
            input: `(define fact 
            (lambda (n) 
                (if (<= n 1) 1 (* n (fact (- n 1))))))`,
            want: "0",
        },
        {
            input: "(fact 10)",
            want:  "3628800",
        },
        {
            input: "(define twice (lambda (x) (* 2 x)))",
            want:  "0",
        },
        {
            input: "(twice 5)",
            want:  "10",
        },
        {
            input: "(define repeat (lambda (f) (lambda (x) (f (f x)))))",
            want:  "0",
        },
        {
            input: "((repeat twice) 10)",
            want:  "40",
        },
        {
            input: "((repeat (repeat twice)) 10)",
            want:  "160",
        },
        {
            input: "((repeat (repeat (repeat twice))) 10)",
            want:  "2560",
        },
        {
            input: "((repeat (repeat (repeat (repeat twice)))) 10)",
            want:  "655360",
        },
        {
            input: `((lambda (a b) (cond ((= a 4) 6)
                          ((= b 4) (+ 6 7))
                          (else 25))) 1 4)`,
            want:  "13",
        },
    }{
        p := parse(tt.input)
        e := evalEnv(p, env)
        got := e.String()
        if got != tt.want {
            t.Errorf("%d) got %s want %s", i, got, tt.want)
        }
    }
}
