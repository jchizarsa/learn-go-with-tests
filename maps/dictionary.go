// Things learned:
// - Created a CRUD (Create, Read, Update, Delete) api using Go maps
//		- Created Go maps
//		- Implemented the following features: Search, Add, Update, Delete to manipulate items in a map
// 		- Practiced more error handling (Creating errors as constants and wrote error wrappers)

// NEVER DO THIS:
// var m map[string]string // This initializes a nil map, which causes runtime panic when there are attempts to add to it.
//
// Always do these instead:
// Initialize as empty - var dictionary = map[string]string{}
// Use the "make" keyword - var dictionary = make(map[string]string)

package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for :(")
	ErrWordExists       = DictionaryErr("this word already exists silly :)")
	ErrWordDoesNotExist = DictionaryErr("that's not in the dictionary :/") // Although ErrNotFound could be reused, having a discrete reason for the error is often beneficial.
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err { // Introducing switch statements in Go!
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word) // Built-in Go function to delete that works on maps
	default:
		return err
	}

	return nil
}
