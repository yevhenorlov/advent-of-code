; --- Day 3: Perfectly Spherical Houses in a Vacuum ---

; Santa is delivering presents to an infinite two-dimensional grid of houses.

; He begins by delivering a present to the house at his starting location, and then an elf at the North Pole calls him via radio and tells him where to move next. Moves are always exactly one house to the north (^), south (v), east (>), or west (<). After each move, he delivers another present to the house at his new location.

; However, the elf back at the north pole has had a little too much eggnog, and so his directions are a little off, and Santa ends up visiting some houses more than once. How many houses receive at least one present?

; For example:

;     > delivers presents to 2 houses: one at the starting location, and one to the east.
;     ^>v< delivers presents to 4 houses in a square, including twice to the house at his starting/ending location.
;     ^v^v^v^v^v delivers a bunch of presents to some very lucky children at only 2 houses.

(ns advent_of_code_2015.03
  (:gen-class)
  (:require [clojure.string :as s]))
(def input (slurp "input03.txt"))

(defn get-movement-modifiers
  "get x and y functions for updated position based on direction"
  [direction]
  (cond
    (s/starts-with? direction "^") [identity inc]
    (s/starts-with? direction "v") [identity dec]
    (s/starts-with? direction "<") [dec identity]
    (s/starts-with? direction ">") [inc identity]
    :else [identity identity]))

(defn move
  "apply movement functions to coords and get new coords"
  [[x-modifier y-modifier] [x y]]
  [(x-modifier x) (y-modifier y)])

(defn get-locations
  "convert list of directions to log of coordinates traversed"
  [directions]
  (loop
   [[head & tail] directions
    coords [0 0]
    result []]
    (if (nil? head)
      result
      (recur tail
             (move (get-movement-modifiers head) coords)
             (conj result coords)))))

(count (set (get-locations input)))

; --- Part Two ---

; The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

; Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

; This year, how many houses receive at least one present?

; For example:

;     ^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
;     ^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
;     ^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.

(defn split-instructions
  [instructions]
  (loop [[[index el] & tail] (map-indexed vector instructions)
         [santa robo-santa] [[] []]]
    (if (nil? el)
      [santa robo-santa]
      (recur tail
             (if (even? index)
               [(conj santa el) robo-santa]
               [santa (conj robo-santa el)])))))

(let [[santa robo-santa] (split-instructions input)]
  (count (set (into (get-locations santa)
                    (get-locations robo-santa)))))

