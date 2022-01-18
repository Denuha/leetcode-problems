package problems

/*
	https://leetcode.com/problems/powx-n/submissions/
	http://mech.math.msu.su/~shvetz/54/inf/perl-problems/chFastPower_sIdeas.xhtml
*/

func myPow(x float64, n int) float64 {
	var count float64 = 1
	if n == 0 {
		return 1
	}

	if n < 0 {
		n *= -1
		x = 1 / x
	}

	for n > 0 {
		if n%2 == 0 {
			n /= 2
			x *= x
		} else {
			n -= 1
			count *= x
		}
	}

	return count
}
