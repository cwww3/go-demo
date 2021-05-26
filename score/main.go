package main

type Game struct {
	bestScore int
	scores    chan int
	sign      chan struct{}
}

func (g *Game) run() {
	for score := range g.scores {
		if score > 100 {
			break
		}
		if score > g.bestScore {
			g.bestScore = score
		}
	}
}

func NewGame() *Game {
	g := &Game{
		scores: make(chan int),
		sign:   make(chan struct{}),
	}
	go g.run()
	return g
}

type Player interface {
	NextScore() (score int, err error)
}

func (g *Game) HandelPlayer(p Player) error {
	for {
		select {
		case <-g.sign:
			return nil
		default:
			if score, err := p.NextScore(); err != nil {
				return err
			} else {
				g.scores <- score
			}
		}
	}
}
