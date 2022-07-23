package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default :
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.

/*
If you have a pair of aces you must always split them.
If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
If your cards sum up to a value within the range [17, 20] you should always stand.
If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
If your cards sum up to 11 or lower you should always hit.
*/
func FirstTurn(card1, card2, dealerCard string) string {
	o := ParseCard(card1)
	t := ParseCard(card2)
	d := ParseCard(dealerCard)
	switch {
	case o + t == 22:
		return "P"
	case o + t == 21:
		if d == 10 || d == 11{
			return "S"
		}
		return "W"
	case o + t >= 17 && o+t <= 20:
		return "S"
	case o+t >=12 && o+t <=16:
		if d >= 7{
			return "H"
		}
		return "S"
	default:
		return "H"
}
}
