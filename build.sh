#/bin/sh

THRIFT_BINARY="/usr/local/bin/thrift"
PARAMS="-gen go:package_prefix=github.com/shinji310/evernote-sdk-go/,thrift_import=github.com/apache/thrift/lib/go/thrift -out ."
THRIFT_FILES_DIR="../../evernote/evernote-thrift/src"

$THRIFT_BINARY $PARAMS $THRIFT_FILES_DIR/Errors.thrift
$THRIFT_BINARY $PARAMS $THRIFT_FILES_DIR/Limits.thrift
$THRIFT_BINARY $PARAMS $THRIFT_FILES_DIR/NoteStore.thrift
$THRIFT_BINARY $PARAMS $THRIFT_FILES_DIR/Types.thrift
$THRIFT_BINARY $PARAMS $THRIFT_FILES_DIR/UserStore.thrift
