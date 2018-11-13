; Define the double procedure
(define double (lambda (x) (+ x x)))

; Define the square procedure
(define square (lambda (x) (* x x)))

; Call square on the result of double
(square (double 3))