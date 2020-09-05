package itry

//Block Try Catch Block
type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

//Exception Interface for an exception
type Exception interface{}

//Throw Once the error is thrown
func Throw(up Exception) {
	panic(up)
}

//Do Function to be called after defining try catch
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
