package sender

import "testing"

func TestChunkSizeMatchesRsyncProtocol(t *testing.T) {
	// Rsync receivers with the c44c90e hardening change reject literal tokens
	// larger than the protocol's CHUNK_SIZE, including on the fast path for a
	// new file that has no block signatures.
	if got, want := chunkSize, 32*1024; got != want {
		t.Fatalf("chunkSize = %d, want %d", got, want)
	}
}
