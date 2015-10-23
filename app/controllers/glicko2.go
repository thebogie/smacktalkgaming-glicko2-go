package controllers

import (
	"github.com/revel/revel"
	"math"
	"mitchgottlieb.com/smacktalkgaming/app/models"
	"strconv"
)

const DEFAULT_RATING float64 = 1500.0
const DEFAULT_DEVIATION float64 = 350
const DEFAULT_VOLATILITY float64 = 0.06
const DEFAULT_TAU float64 = 0.75
const MULTIPLIER float64 = 173.7178
const CONVERGENCE_TOLERANCE float64 = 0.000001

const POINTS_FOR_WIN float64 = 1.0
const POINTS_FOR_LOSS float64 = 0.0
const POINTS_FOR_DRAW float64 = 0.5

/**** RATING CALCULATOR ****/
type RatingCalculator struct {
	tau               float64
	defaultVolatility float64
}

func (rc *RatingCalculator) convertRatingToGlicko2Scale(rating float64) float64 {
	return ((rating - DEFAULT_RATING) / MULTIPLIER)
}

func (rc *RatingCalculator) convertRatingToOriginalGlickoScale(rating float64) float64 {
	return ((rating * MULTIPLIER) + DEFAULT_RATING)
}

func (rc *RatingCalculator) convertRatingDeviationToGlicko2Scale(ratingDeviation float64) float64 {
	return (ratingDeviation / MULTIPLIER)

}

func (rc *RatingCalculator) convertRatingDeviationToOriginalGlickoScale(ratingDeviation float64) float64 {
	return (ratingDeviation * MULTIPLIER)

}

func (rc *RatingCalculator) findF(x float64, delta float64, phi float64, v float64, a float64, tau float64) float64 {
	return (math.Exp(x) * (math.Pow(delta, 2) - math.Pow(phi, 2) - v - math.Exp(x)) / (2.0 * math.Pow(math.Pow(phi, 2)+v+math.Exp(x), 2))) - ((x - a) / math.Pow(tau, 2))
}

func (rc *RatingCalculator) findG(deviation float64) float64 {
	return 1.0 / (math.Sqrt(1.0 + (3.0 * math.Pow(deviation, 2) / math.Pow(math.Pi, 2))))
}

func (rc *RatingCalculator) findE(playerRating float64, opponentRating float64, opponentDeviation float64) float64 {

	return 1.0 / (1.0 + math.Exp(-1.0*rc.findG(opponentDeviation)*(playerRating-opponentRating)))
}

func (rc *RatingCalculator) outcomeBasedRating(player Rating, results []Result) float64 {
	outcomeBasedRating := 0.0

	for _, result := range results {
		opponent := result.getOpponent(player)
		outcomeBasedRating = outcomeBasedRating + (rc.findG(opponent.getGlicko2RatingDeviation()) * (result.getScore(player) - rc.findE(player.getGlicko2Rating(), opponent.getGlicko2Rating(), opponent.getGlicko2RatingDeviation())))
	}

	return outcomeBasedRating
}

func (rc *RatingCalculator) findV(player Rating, results []Result) float64 {

	v := 0.0
	for _, result := range results {
		opponent := result.getOpponent(player)
		v = v + ((math.Pow(rc.findG(opponent.getGlicko2RatingDeviation()), 2)) *
			rc.findE(player.getGlicko2Rating(), opponent.getGlicko2Rating(), opponent.getGlicko2RatingDeviation()) *
			(1.0 - rc.findE(player.getGlicko2Rating(), opponent.getGlicko2Rating(), opponent.getGlicko2RatingDeviation())))
	}

	return math.Pow(v, -1)

}

func (rc *RatingCalculator) findDelta(player Rating, results []Result) float64 {
	return rc.findV(player, results) * rc.outcomeBasedRating(player, results)
}

func (rc *RatingCalculator) calculateNewRD(phi float64, sigma float64) float64 {
	return math.Sqrt(math.Pow(phi, 2) + math.Pow(sigma, 2))
}

/*+++ END OF RATINGCALCULATOR OBJECT +++*/

/**** RATING OBJECT ****/
type Rating struct {
	playerName             string
	playerUUID             string
	uid                    string
	rating                 float64
	ratingDeviation        float64
	volatility             float64
	numberOfResults        int
	workingRating          float64
	workingRatingDeviation float64
	workingVolatility      float64
	ratingcalc             RatingCalculator
}

func (r *Rating) setup(uuid string, ratingSystem RatingCalculator, initRating float64, initRatingDeviation float64, initVolatility float64) {

	r.uid = uuid
	r.rating = initRating
	r.ratingDeviation = initRatingDeviation
	r.volatility = initVolatility
}

func (r *Rating) getRating() float64 {
	return r.rating
}
func (r *Rating) setRating(rating float64) {
	r.rating = rating
}

func (r *Rating) getVolatility() float64 {
	return r.volatility
}
func (r *Rating) setVolatility(volatility float64) {
	r.volatility = volatility
}

func (r *Rating) getRatingDeviation() float64 {
	return r.ratingDeviation
}
func (r *Rating) setRatingDeviation(ratingDeviation float64) {
	r.ratingDeviation = ratingDeviation
}

func (r *Rating) getGlicko2Rating() float64 {

	return r.ratingcalc.convertRatingToGlicko2Scale(r.rating)

}

func (r *Rating) getGlicko2RatingDeviation() float64 {
	return r.ratingcalc.convertRatingDeviationToGlicko2Scale(r.ratingDeviation)
}

/**
 * Returns a formatted rating for inspection
 *
 * @return {ratingUid} / {ratingDeviation} / {volatility} / {numberOfResults}
 */

/*
func (r *Rating) toString() string {
	return r.uid + " / " +
		r.rating + " / " +
		r.ratingDeviation + " / " +
		r.volatility + " / " +
		r.numberOfResults
}
*/

func (r *Rating) getNumberOfResults() int {
	return r.numberOfResults
}

func (r *Rating) incrementNumberOfResults(increment int) {
	r.numberOfResults = r.numberOfResults + increment
}

func (r *Rating) getUid() string {
	return r.uid
}

func (r *Rating) convertToFloats(node models.Glicko2, playeruuid string, playername string) {

	r.playerName = playername
	r.playerUUID = playeruuid
	r.uid = node.UUID
	r.numberOfResults, _ = strconv.Atoi(node.NumResults)
	r.rating, _ = strconv.ParseFloat(node.Rating, 64)
	r.ratingDeviation, _ = strconv.ParseFloat(node.RatingDeviation, 64)
	r.volatility, _ = strconv.ParseFloat(node.Volatility, 64)
}

func (r *Rating) finaliseRating() (node models.Glicko2) {
	revel.TRACE.Println("ITEMS", r)

	r.rating = r.ratingcalc.convertRatingToOriginalGlickoScale(r.workingRating)
	r.ratingDeviation = r.ratingcalc.convertRatingDeviationToOriginalGlickoScale(r.workingRatingDeviation)
	r.volatility = r.workingVolatility

	node.UUID = r.uid
	node.Rating = strconv.FormatFloat(r.rating, 'f', 10, 64)
	node.RatingDeviation = strconv.FormatFloat(r.ratingDeviation, 'f', 10, 64)
	node.Volatility = strconv.FormatFloat(r.volatility, 'f', 10, 64)

	r.workingVolatility = 0
	r.workingRating = 0
	r.workingRatingDeviation = 0

	revel.TRACE.Println("ITEMS AFTER", r)

	return node

}

/*+++ END OF RATING OBJECT +++*/

/**** RESULT OBJECT ****/
type Result struct {
	isDraw bool
	winner Rating
	loser  Rating
}

func (res *Result) init(winner Rating, loser Rating) {
	if !res.validPlayers(winner, loser) {
		//TODO: throw golang error
	}
	res.winner = winner
	res.loser = loser

}

func (res *Result) getScore(player Rating) float64 {
	var score float64

	if res.winner.playerUUID == player.playerUUID {
		score = POINTS_FOR_WIN
	} else if res.loser.playerUUID == player.playerUUID {
		score = POINTS_FOR_LOSS
	} else {
		//throw new IllegalArgumentException("Player " + player.getUid() + " did not participate in match");
	}

	if res.isDraw {
		score = POINTS_FOR_DRAW
	}

	return score
}

func (res *Result) validPlayers(player1 Rating, player2 Rating) bool {
	if player1.playerUUID == (player2.playerUUID) {
		return false
	} else {
		return true
	}
}
func (res *Result) participated(player Rating) bool {
	if res.winner.playerUUID == player.playerUUID || res.loser.playerUUID == player.playerUUID {
		return true
	} else {
		return false
	}
}

func (res *Result) getOpponent(player Rating) Rating {
	opponent := Rating{}

	if res.winner.playerUUID == player.playerUUID {
		opponent = res.loser
	} else if res.loser.playerUUID == player.playerUUID {
		opponent = res.winner
	} else {
		//TODO: Throw golang error
		//throw new IllegalArgumentException("Player " + player.getUid() + " did not participate in match");
	}

	return opponent
}

func (res *Result) getWinner() Rating {
	return res.winner
}

func (res *Result) getLoser() Rating {
	return res.loser
}

/*+++ END OF RESULT OBJECT +++*/

/**** RATINGPERIODRESULTS  OBJECT ****/
type RatingPeriodResults struct {
	results []Result
}

func (rp *RatingPeriodResults) addResult(winner Rating, loser Rating) {

	result := new(Result)
	result.init(winner, loser)

	rp.results = append(rp.results, *result)
}

/*+++ END OF RATINGPERIODRESULTS OBJECT +++*/

type Glicko2Rating struct {
	Application
}

//TODO: need to optimize. Lots of loops becuase trying to keep the same language
// as the java version to get glicko2 working
func afterEventRankingUpdate(UUIDEvt string) {
	var eventresults = new(models.QueryObj).GetCompetitorsByEvent(UUIDEvt)
	var rc = RatingCalculator{tau: DEFAULT_TAU, defaultVolatility: DEFAULT_VOLATILITY}

	finalize := []Rating{}

	revel.TRACE.Println("eventresults", eventresults)

	for index, c := range eventresults {

		mainplayer := Rating{}
		mainplayer.convertToFloats(new(models.QueryObj).GetPlayerGlicko2Rating(c.Player.UUID), c.Player.UUID, c.Player.Firstname)
		var resultperiod = RatingPeriodResults{}
		mainplayer.ratingcalc = rc

		for i := 0; i < len(eventresults); i++ {

			if !(c.Player.UUID == eventresults[i].Player.UUID) {
				revel.TRACE.Println("Compare ", c.Player.Firstname, "to", eventresults[i].Player.Firstname)
				revel.TRACE.Println("with eventresults ", eventresults[index].Result.Place, ":", eventresults[index].Result.Result, "to", eventresults[i].Result.Place, ":", eventresults[i].Result.Result)

				opponent := Rating{}
				opponent.ratingcalc = rc
				opponent.convertToFloats(new(models.QueryObj).GetPlayerGlicko2Rating(eventresults[i].Player.UUID), eventresults[i].Player.UUID, eventresults[i].Player.Firstname)

				loser := Rating{}
				winner := Rating{}
				loser.ratingcalc = rc
				winner.ratingcalc = rc

				mainplayerPlace, _ := strconv.Atoi(c.Result.Place)
				otherplayerPlace, _ := strconv.Atoi(eventresults[i].Result.Place)
				if mainplayerPlace < otherplayerPlace {

					winner = mainplayer
					loser = opponent

					revel.TRACE.Println("WINNER = MAINPLAYER", c.Player.Firstname)
					revel.TRACE.Println("LOSER = OTHERPLAYER", opponent.playerName)

				} else {
					winner = opponent
					loser = mainplayer

					revel.TRACE.Println("WINNER = OHTERPLAYER", opponent.playerName)
					revel.TRACE.Println("LOSER = MAINPLAYER", c.Player.Firstname)

				}
				resultperiod.addResult(winner, loser)
			}

		}

		//calculate
		phi := mainplayer.getGlicko2RatingDeviation()
		sigma := mainplayer.volatility
		a := math.Log(math.Pow(sigma, 2))
		delta := rc.findDelta(mainplayer, resultperiod.results)
		v := rc.findV(mainplayer, resultperiod.results)

		// step 5.2 - set the initial values of the iterative algorithm to come in step 5.4
		A := a
		B := 0.0
		if math.Pow(delta, 2) > math.Pow(phi, 2)+v {
			B = math.Log(math.Pow(delta, 2) - math.Pow(phi, 2) - v)
		} else {
			k := 1.0
			B = a - (k * math.Abs(DEFAULT_TAU))

			for rc.findF(B, delta, phi, v, a, DEFAULT_TAU) < 0 {
				k = k + 1
				B = a - (k * math.Abs(DEFAULT_TAU))
			}
		}
		// step 5.3
		fA := rc.findF(A, delta, phi, v, a, DEFAULT_TAU)
		fB := rc.findF(B, delta, phi, v, a, DEFAULT_TAU)

		// step 5.4
		for math.Abs(B-A) > CONVERGENCE_TOLERANCE {
			C := A + (((A - B) * fA) / (fB - fA))
			fC := rc.findF(C, delta, phi, v, a, DEFAULT_TAU)

			if fC*fB < 0 {
				A = B
				fA = fB
			} else {
				fA = fA / 2.0
			}

			B = C
			fB = fC
		}

		newSigma := math.Exp(A / 2.0)

		mainplayer.workingVolatility = newSigma

		// Step 6
		phiStar := rc.calculateNewRD(phi, newSigma)

		// Step 7
		newPhi := 1.0 / math.Sqrt((1.0/math.Pow(phiStar, 2))+(1.0/v))

		// note that the newly calculated rating values are stored in a "working" area in the Rating object
		// this avoids us attempting to calculate subsequent participants' ratings against a moving target
		mainplayer.workingRating = mainplayer.getGlicko2Rating() + (math.Pow(newPhi, 2) * rc.outcomeBasedRating(mainplayer, resultperiod.results))
		mainplayer.workingRatingDeviation = newPhi
		mainplayer.incrementNumberOfResults(len(resultperiod.results))

		finalize = append(finalize, mainplayer)

	}

	updatedratings := []models.Glicko2{}
	for _, final := range finalize {

		updatedratings = append(updatedratings, final.finaliseRating())

	}

	revel.TRACE.Println("final results", updatedratings)

}
