;; Define the double procedure
(define double (lambda (x) (+ x x)))

;; Define the square procedure
(define square (lambda (x) (* x x)))

(define five 5)

;; Call square on the result of double
(square (double (+ five 3)))

;; Booleans
(and #f #t)
(or #f #t)

(/ 4 3)

(+ 1 2 3)             ;; => 6 ; this S-expression evaluates to a function call
(quote (+ 1 2 3))     ;; => (+ 1 2 3) ; evaluates to a list


(car (list 1 2 3))    ;; Constructs a list from the pair and returns its head
(car (quote (1 2 3))) ;; Quotes the list instead of evaluating it and returns the list head
