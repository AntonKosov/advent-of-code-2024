package math

type Vector2[T Numbers] struct {
	X T
	Y T
}

func NewVector2[T Numbers](x, y T) Vector2[T] {
	return Vector2[T]{X: x, Y: y}
}

func (v Vector2[T]) Add(av Vector2[T]) Vector2[T] {
	return NewVector2(v.X+av.X, v.Y+av.Y)
}

func (v Vector2[T]) Sub(av Vector2[T]) Vector2[T] {
	return NewVector2(v.X-av.X, v.Y-av.Y)
}
