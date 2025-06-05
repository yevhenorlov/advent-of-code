; --- Day 4: The Ideal Stocking Stuffer ---

; Santa needs help mining some AdventCoins (very similar to bitcoins) to use as gifts for all the economically forward-thinking little girls and boys.

; To do this, he needs to find MD5 hashes which, in hexadecimal, start with at least five zeroes. The input to the MD5 hash is some secret key (your puzzle input, given below) followed by a number in decimal. To mine AdventCoins, you must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...) that produces such a hash.

; For example:

;     If your secret key is abcdef, the answer is 609043, because the MD5 hash of abcdef609043 starts with five zeroes (000001dbbfa...), and it is the lowest such number to do so.
;     If your secret key is pqrstuv, the lowest number it combines with to make an MD5 hash starting with five zeroes is 1048970; that is, the MD5 hash of pqrstuv1048970 looks like 000006136ef....

(import 'java.security.MessageDigest
        'java.math.BigInteger)
(ns advent_of_code_2015.04
  (:gen-class))
(def input "iwrupvqb")

(defn md5 [s]
  (let [algorithm (MessageDigest/getInstance "MD5")
        raw (.digest algorithm (.getBytes s))]
    (format "%032x" (BigInteger. 1 raw))))

(defn gen
  [x]
  (md5 (str input x)))

(defn fits-criteria? 
  [x]
  (clojure.string/starts-with? (gen x) "00000")) ;; or "000000"

(first (drop-while (complement #(fits-criteria? %)) (range)))

; answers:
; (gen "346386")
; (gen "9958218")
