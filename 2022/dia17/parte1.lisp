

;; hashmap contendo posições de todos os blocos de todas as peças. todas as
;; colisões são calculadas aqui
(defparameter *chamber* (make-hash-table :test #'equalp))

(defparameter *gas* (uiop:read-file-string "hotgas.txt"))

(defparameter *gas-len* (length *gas*))

(defparameter *gas-current* 0)

;; posições relativas de cada bloco de cada peça.
;; ultimo bloco é o bloco mais inferior
;; primeiro bloco é o mais a esquerda
(defparameter *shapes*
  (list '((0 . 0) (1 . 0) (2 . 0) (3 . 0))
        '((0 . 1) (1 . 1) (2 . 1) (1 . 2) (1 . 0))
        '((0 . 0) (1 . 0) (2 . 2) (2 . 1) (2 . 0))
        '((0 . 3) (0 . 2) (0 . 1) (0 . 0))
        '((0 . 1) (1 . 1) (0 . 0) (1 . 0))))

(defun x (shape)
  (car shape))

(defun y (shape)
  (cdr shape))

(defun current-rock (x)
  ;; retornar ao inicio quando todas as peças forem usadas
  (nth (mod x 5) *shapes*))

(defun current-flow () ;; feio
  (prog1 (aref *gas* (mod *gas-current* *gas-len*))
    (incf *gas-current*)))

(defun highest (shape)
  (1+ (loop for x in shape
        maximize (y x))))

;; move peça mas não adiciona a *chamber*
(defun transform-shape (shape x y)
  (loop with transform = shape
        for j in transform
	collect (cons (+ (x j) x)
		      (+ (y j) y))))

(defun move-down (shape)
  (transform-shape shape 0 -1))

;; verifica se o bloco sobrepoe algum outro bloco ja existente
(defun collided-p (shape)
  (loop for i in shape
	thereis (< (cdr i) 0)
        thereis (gethash i *chamber*)))

;; posiciona peça definitivamente em *chamber*
(defun place-shape (shape)
  (loop for x in shape
        do (setf (gethash x *chamber*) t)))

(defun chamber-fit-p (shape)
  (loop for i in shape
	never (< (x i) 0)
	never (> (x i) 6)))

(defun gas-direction (gas shape)
  (let ((moved (if (equalp gas #\<)
		      (transform-shape shape -1 0)
		      (transform-shape shape  1 0))))
    (if (and (chamber-fit-p moved)
	     ;; possivel que haja blocos ja nessa posição
	     (not (collided-p moved)))
	moved
	shape)))

(defun simulate-rock (highest rock)
  ;; ajustar posição da peça para que fique no lugar correto
  (loop with current = (transform-shape rock 2 (+ 3 highest))
	with down = nil
	;; mover peça na direção necessaria
	do (setf current (gas-direction (current-flow) current)
		 ;; peça caso movida para baixo
		 down (move-down current))
	;; verifica se há espaço para continuar a descida
        if (collided-p down)
          do (place-shape current)
	  and do (return (highest current))
	;; mover peça para baixo ja que colisão não é possivel
	else do (setf current down)))

(setf *gas-current* 0)
(setf *chamber* (make-hash-table :test #'equalp))

(loop for x from 0 to 3000
      ;; do (format t "highest: ~a~%" highest)
      maximize (simulate-rock highest (current-rock x)) into highest
      finally (print highest))
