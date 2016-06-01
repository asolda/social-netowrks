package metrics

import (
	"github.com/asolda/sim/framework"
	"github.com/dcadenas/pagerank"
)

func PageRank(graph framework.Graph) []float64 {

	g := pagerank.New()

	ranks := make([]float64, len(graph.GetAgents()))

	for i := 0; i < len(graph.GetAgents()); i++ {
		ranks[i] = 0
	}

	for _, agent := range graph.GetAgents() {
		for _, edge := range agent.GetConnections() {
			g.Link(agent.GetID(), edge.To.GetID())
		}
	}

	prob := 0.85        // The bigger the number, less probability we have to teleport to some random link
	tolerance := 0.0001 // the smaller the number, the more exact the result will be but more CPU cycles will be needed

	g.Rank(prob, tolerance, func(identifier int, rank float64) {
		ranks[identifier] = rank
	})

	return ranks
}
