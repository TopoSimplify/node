package node

import (
	"testing"
	"simplex/pln"
	"simplex/rng"
	"github.com/intdxdt/geom"
	"github.com/franela/goblin"
)

var fn = hullGeom

type testD struct {
	k   int
	bln bool
	wkt string
}

func TestCollapsible(t *testing.T) {
	var g = goblin.Goblin(t)
	var lnwkts = []testD{
		{21, false,
			"LINESTRING ( 810 540, 790 570, 800 580, 820 580, 860 570, 880 600, 870 610, 850 610, 800 610, 810 650, 890 640, 900 640, 920 600, 930 580, 930 540, 920 500, 880 490, 860 520, 810 510, 750 520, 780 460, 730 410, 830 440, 890 410, 940 450, 970 500, 1040 510, 1050 570, 1080 620, 1040 660, 1020 720, 950 720, 840 680, 760 690, 690 720, 710 640, 630 620 )"},
		{4, true,
			"LINESTRING (330.5770802442433 581.2490753170464, 289.3831572774428 587.7108671549759, 256.2664741080541 582.8645232765288, 197.441310580554 568.7784526288414, 194.9264255896295 552.0343728591217, 202.11368678741684 562.5981327919068, 211.82048369391282 548.4791554733671, 216.23266410595645 527.5948348563605, 218.58582699237974 516.711456506653, 185.64154658245394 504.3573513529308, 176.52304039756376 493.17982764242026, 172.1108599855201 481.4140132103039, 172.69915070712594 461.706274036509, 157.69773730617757 464.3535822837352, 137.9899981323827 464.941873005341)"},
		{4, true,
			"LINESTRING ( 330.5770802442433 581.2490753170464, 289.3831572774428 587.7108671549759, 256.2664741080541 582.8645232765288, 197.441310580554 568.7784526288414, 194.9264255896295 552.0343728591217, 202.11368678741684 562.5981327919068, 211.74199479333123 560.2115030178386, 227.0827680621383 567.7253511494993, 236.3364918395976 564.4237400821951, 249.31123545163425 568.9776591714426, 241.34572392737135 584.7737454387759, 235.22277020477063 586.8230484841366, 227.70892207311002 596.8415126596841, 222.69968998533628 614.3738249668921, 180 620 )"},
		{4, true,
			"LINESTRING ( 330.5770802442433 581.2490753170464, 289.3831572774428 587.7108671549759, 256.2664741080541 582.8645232765288, 197.441310580554 568.7784526288414, 194.9264255896295 552.0343728591217, 185.35048559505015 556.3488518933807, 195.5654156879096 545.0472271097915, 227.0827680621383 567.7253511494993, 236.3364918395976 564.4237400821951, 237.11918435331225 572.5637422248274, 241.34572392737135 584.7737454387759, 209.2553308650708 583.2083604113466, 207.2203303294127 590.2525930347784, 204.40263728003995 601.9929807404982, 189.2184025139758 610.2895213858734 )"},
		{4, false,
			"LINESTRING ( 230.0195128864249 579.5695574923628, 217.3543939499459 574.8349335908753, 210.8442860854006 568.916653714016, 197.441310580554 568.7784526288414, 194.9264255896295 552.0343728591217, 180 560, 190.34928117240688 543.9605324190618, 210 550, 236.3364918395976 564.4237400821951, 237.11918435331225 572.5637422248274, 241.34572392737135 584.7737454387759, 209.2553308650708 583.2083604113466, 207.2203303294127 590.2525930347784, 204.40263728003995 601.9929807404982, 189.2184025139758 610.2895213858734 )"},
		{7, false,
			"LINESTRING ( 208 576, 212 568, 206 572, 208 562, 202 568, 196 564, 200 560, 194 554, 202 556, 214 556, 220 560, 220 568, 226 572, 220 580, 212 584, 210 590, 204 586, 200 596, 192 588, 196 588, 202 582 )"},
		{7, false,
			"LINESTRING ( 208 576, 212 568, 206 572, 208 562, 202 568, 196 564, 200 560, 194 554, 202 556, 214 556, 220 560, 220 568, 226 572, 220 580, 212 584, 210 590, 208 584, 204 588, 202 582, 204 584, 208 580 )"},
	}

	g.Describe("hull collapse", func() {
		g.It("should test hull collapsibility", func() {
			for _, o := range lnwkts {
				k, bln, wkt := o.k, o.bln, o.wkt
				var coords = geom.NewLineStringFromWKT(wkt).Coordinates()
				var poly = pln.New(coords)
				var n = len(coords) - 1
				var rng_a , rng_b = rng.NewRange(0, k), rng.NewRange(k, n)
				var ha, hb = New(poly.SubCoordinates(rng_a), rng_a, fn), New(poly.SubCoordinates(rng_b), rng_b, fn)
				g.Assert(hb.Collapsible(ha)).Equal(bln)
			}
		})

		g.It("should test collapsible of contiguous and non contiguous", func() {
			wkt := "LINESTRING ( 467.082432820504 469.7831661127625, 480.006016496363 438.2819309028562, 505.85318384808096 402.74207579424393, 587.433305801941 375.2794604830436, 624.5886088700355 382.54897630071423, 642.3585364243417 412.4347635511382, 643.4709527477502 439.116190719288, 605.3439380779691 452.77166096041014, 574.5798901784301 466.2701717734732, 531.700351199799 481.0913018291391, 523.623111402387 491.59171356577457, 520.3922154834223 515.8234329580102, 521.1999394631636 540.0551523502459, 530.8926272200578 545.7092202084342, 581.1879051657643 561.1022524449884, 599.9685326993344 539.6722729770875, 620.4120431716209 520.2642150257254, 632.0270408479774 503.31259679536714, 643.4709527477502 439.116190719288, 653.6666721407183 466.55227019379777, 656.8975680596831 497.2457814239629, 644.7817083635653 540.862876329987, 593.0873736601293 582.0567992967876, 551.8934506933286 597.4035549118702 )"
			coords := geom.NewLineStringFromWKT(wkt).Coordinates()
			k1, k2, k3, k4 := 6, 10, 14, 18
			n := len(coords) - 1
			polyline := pln.New(coords)

			h1 := New(polyline.SubCoordinates(rng.NewRange(0, k1)), rng.NewRange(0, k1) , fn)
			h2 := New(polyline.SubCoordinates(rng.NewRange(k1, k2) ), rng.NewRange(k1, k2)   , fn)
			h3 := New(polyline.SubCoordinates(rng.NewRange(k2, k3) ), rng.NewRange(k2, k3)   , fn)
			h4 := New(polyline.SubCoordinates(rng.NewRange(k3, k4) ), rng.NewRange(k3, k4)   , fn)
			h5 := New(polyline.SubCoordinates(rng.NewRange(k4, n)  ), rng.NewRange(k4, n)    , fn)

			hulls := [][2]*Node{{h1, h4}, {h1, h2}, {h1, h5}, {h2, h3}, {h2, h4}, {h2, h5}}
			for _, o := range hulls {
				ha, hb := o[0], o[1]
				g.Assert(hb.Collapsible(ha)).IsTrue()
			}

			ha, hb := h4, h5
			g.Assert(hb.Collapsible(ha)).IsFalse()
			//if not contiguous should be eql
			g.Assert(h5.Collapsible(h3)).IsTrue()

		})
	})
}
