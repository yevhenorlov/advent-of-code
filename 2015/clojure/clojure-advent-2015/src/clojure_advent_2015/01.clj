; https://adventofcode.com/2015/day/1

; --- Day 1: Not Quite Lisp ---

; Santa was hoping for a white Christmas, but his weather machine's "snow" function is powered by stars, and he's fresh out! To save Christmas, he needs you to collect fifty stars by December 25th.

; Collect stars by helping Santa solve puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

; Here's an easy puzzle to warm you up.

; Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character at a time.

; An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.

; The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.

; For example:

;     (()) and ()() both result in floor 0.
;     ((( and (()(()( both result in floor 3.
;     ))((((( also results in floor 3.
;     ()) and ))( both result in floor -1 (the first basement level).
;     ))) and )())()) both result in floor -3.

; To what floor do the instructions take Santa?

(ns clojure-advent-2015.01
  (:gen-class)
  (:require [clojure.string :as string]))
(def input (slurp "input01.txt"))

; ---
; first naive attempt, didn't work. 
; keeping as a reminder that recursion is tricky -
; this one fails with a stack overflow error)

; (defn parse
;   ([instructions]
;    (parse instructions 0))
;   ([instructions counter]
;    (cond
;      (nil? (first instructions))
;      counter
;      (string/starts-with? (first instructions) "(")
;      (parse (rest instructions) (inc counter))
;      (string/starts-with? (first instructions) ")")
;      (parse (rest instructions) (dec counter))
;      :else counter)))

; ---
; This works:

(defn count-floor
  ([instructions]
   (reduce
    (fn [acc el]
      (cond
        (string/starts-with? el "(")
        (inc acc)
        (string/starts-with? el ")")
        (dec acc)
        :else acc))
    0
    instructions)))

(count-floor input)

; ---
; Shorter (source: https://ahxxm.com/161.moew/)

(let [f (frequencies input)]
  (- (f \() (f \))))

; --- Part Two ---

; Now, given the same instructions, find the position of the first character that causes him to enter the basement (floor -1). The first character in the instructions has position 1, the second character has position 2, and so on.

; For example:

;     ) causes him to enter the basement at character position 1.
;     ()()) causes him to enter the basement at character position 5.

; What is the position of the character that causes Santa to first enter the basement?

(defn count-basement-entrance
  [instructions]
  (loop
   [[[index el] & remaining] (map-indexed vector instructions)
    floor 0]
    (let [inc-value (if (string/starts-with? el "(") 1 -1)]
      (if (= -1 floor)
        index
        (recur remaining (+ floor inc-value))))))

(count-basement-entrance input)
