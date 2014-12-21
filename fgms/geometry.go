package fgms

import (
	"math"
	"strconv"
)

// SimGear constants
const (
	SG_180                = 180.0
	SG_PI                 = 3.1415926535
	SG_RADIANS_TO_DEGREES = (SG_180 / SG_PI)
	SG_DEGREES_TO_RADIANS = (SG_PI / SG_180)
	SG_FEET_TO_METER      = 0.3048
	SGD_PI_2              = 1.57079632679489661923
)

/*
High-precision versions of the above produced with an arbitrary
precision calculator (the compiler might lose a few bits in the FPU
operations).  These are specified to 81 bits of mantissa, which is
higher than any FPU known to me:
*/
const (
	SQUASH  = 0.9966471893352525192801545
	STRETCH = 1.0033640898209764189003079
	POLRAD  = 6356752.3142451794975639668
)

const (
	// Radians To Nautical Miles
	SG_RAD_TO_NM = 3437.7467707849392526

	// Nautical Miles in a Meter
	SG_NM_TO_METER = 1852.0000

	// Meters to Feet
	SG_METER_TO_FEET = 3.28083989501312335958
)

const (
	X = 0
	Y = 1
	Z = 2
)
const (
	Lat = 0
	Lon = 1
	Alt = 2
)

type Point3D struct {
	X float64
	Y float64
	Z float64
}

//func (me *Point3D) X() float64 { return me._x}
//func (me *Point3D) Y() float64 { return me._y}
//func (me *Point3D) Z() float64 { return me._z}

func (me *Point3D) Lat() float64 { return me.X }
func (me *Point3D) Lon() float64 { return me.Y }
func (me *Point3D) Alt() float64 { return me.Z }

//func (me *Point3D) SetLat(lat float64) { me._x = lat}
//func (me *Point3D) SetLon(lon float64) { me._y = lon}
//func (me *Point3D) SetAlt(alt float64) { me._z = alt}

// returns values to 6 digits with a space between
func (me *Point3D) ToSpacedString() string {
	Message := strconv.FormatFloat(me.X, 'f', 6, 32) + " "
	Message += strconv.FormatFloat(me.Y, 'f', 6, 32) + " "
	Message += strconv.FormatFloat(me.Z, 'f', 6, 32)
	return Message
}

func (me *Point3D) Set(x, y, z float64) {
	me.X = x
	me.Y = y
	me.X = z
}
func (me *Point3D) Clear() {
	me.X = 0
	me.Y = 0
	me.Z = 0
}

func (me *Point3D) Length() float64 {
	//return (sqrt ((m_X * m_X) + (m_Y * m_Y) + (m_Z * m_Z)));
	return math.Sqrt((me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z))
}

func Point3DSubract(p1, p2 Point3D) Point3D {

	return Point3D{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}

}

// Calculate distance of points
func Distance(P1, P2 Point3D) float32 {

	//P = P1 - P2
	var P Point3D
	P = Point3DSubract(P1, P2)

	//return (float)(P.length() / SG_NM_TO_METER);
	return float32(P.Length() / SG_NM_TO_METER)
} // Distance ( const Point3D & P1, const Point3D & P2 )

//-------------------------------------------------------------------

//#define _EQURAD     (6378137.0)
const _EQURAD = 6378137.0

//#define E2 fabs(1 - SQUASH*SQUASH)
var e2 float64 = math.Abs(1 - SQUASH*SQUASH)

//static double ra2 = 1/(_EQURAD*_EQURAD);
var ra2 float64 = 1 / (_EQURAD * _EQURAD)

//static double e2 = E2;
//static double e4 = E2*E2;
var e4 float64 = e2 * e2

/*
 Convert a cartexian XYZ coordinate to a geodetic lat/lon/alt.
   This function is a copy of what's in SimGear,
  simgear/math/SGGeodesy.cxx and fgms http://gitorious.org/fgms/fgms-0-x/blobs/master/src/server/fg_geometry.cxx#line407
*/

func SG_CartToGeod(CartPoint Point3D) Point3D {

	// according to
	// H. Vermeille,
	// Direct transformation from geocentric to geodetic cordinates,
	// Journal of Geodesy (2002) 76:451-454
	//double x = CartPoint[X];
	//double y = CartPoint[Y];
	//double z = CartPoint[Z];
	x := CartPoint.X
	y := CartPoint.Y
	z := CartPoint.Z

	//double XXpYY = x*x+y*y;
	var XXpYY float64 = (x * x) + (y * y)

	//double sqrtXXpYY = sqrt(XXpYY);
	var sqrtXXpYY float64 = math.Sqrt(XXpYY)

	//double p = XXpYY*ra2;
	var p float64 = XXpYY * ra2

	//double q = z*z*(1-e2)*ra2;
	var q float64 = z * z * (1 - e2) * ra2

	//double r = 1/6.0*(p+q-e4);
	var r float64 = 1 / 6.0 * (p + q - e4)

	//double s = e4*p*q/(4*r*r*r);
	var s float64 = e4 * p * q / (4 * r * r * r)

	//double t = pow(1+s+sqrt(s*(2+s)), 1/3.0);
	var t float64 = math.Pow(1+s+math.Sqrt(s*(2+s)), 1/3.0)

	//double u = r*(1+t+1/t);
	var u float64 = r * (1 + t + 1/t)

	//double v = sqrt(u*u+e4*q);
	var v float64 = math.Sqrt(u*u + e4*q)

	//double w = e2*(u+v-q)/(2*v);
	var w float64 = e2 * (u + v - q) / (2 * v)

	//double k = sqrt(u+v+w*w)-w;
	var k float64 = math.Sqrt(u+v+w*w) - w

	//double D = k*sqrtXXpYY/(k+e2);
	var D float64 = k * sqrtXXpYY / (k + e2)

	var GeodPoint Point3D
	//GeodPoint[Lon] = (2*atan2(y, x+sqrtXXpYY)) * SG_RADIANS_TO_DEGREES;
	GeodPoint.Y = (2 * math.Atan2(y, x+sqrtXXpYY)) * SG_RADIANS_TO_DEGREES

	//double sqrtDDpZZ = sqrt(D*D+z*z);
	var sqrtDDpZZ float64 = math.Sqrt(D*D + z*z)
	//GeodPoint[Lat] = (2*atan2(z, D+sqrtDDpZZ)) * SG_RADIANS_TO_DEGREES;
	GeodPoint.X = (2 * math.Atan2(z, D+sqrtDDpZZ)) * SG_RADIANS_TO_DEGREES

	//GeodPoint[Alt] = ((k+e2-1)*sqrtDDpZZ/k) * SG_METER_TO_FEET;
	GeodPoint.Z = ((k + e2 - 1) * sqrtDDpZZ / k) * SG_METER_TO_FEET
	return GeodPoint
} // sgCartToGeod()
