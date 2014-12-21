package hash

type Scheme interface {
  // Hashes a plaintext UTF-8 password using a modular crypt stub. Returns the
  // hashed password in modular crypt format.
  //
  // A modular crypt stub is a prefix of a hash in modular crypt format which
  // expresses all necessary configuration information, such as salt and
  // iteration count. For example, for sha256-crypt, a valid stub would be:
  //
  //     $5$rounds=6000$salt
  //
  // A full modular crypt hash may also be passed as the stub, in which case
  // the hash is ignored.
  Hash(password, stub string) (string, error)

  // Verifies a plaintext UTF-8 password using a modular crypt hash.  Returns
  // an error if the inputs are malformed or the password does not match.
  //
  // The newHash output is ordinarily empty. If it is not empty, it contains an
  // upgraded password hash which should replace the hash which was passed in
  // whereever it is stored.
  Verify(password, hash string) (newHash string, err error)

  // Returns true iff this crypter supports the given stub.
  SupportsStub(stub string) bool

  // Returns true iff this stub needs an update.
  NeedsUpdate(stub string) bool

  // Make a stub with the configured defaults. The salt is generated randomly.
  MakeStub() (string, error)
}

// © 2014 Hugo Landau <hlandau@devever.net>  BSD License
