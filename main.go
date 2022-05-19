package main

/*
#cgo LDFLAGS: -lpam -fPIC
#cgo CFLAGS: -Wall
#include <nss.h>
#include <pwd.h>
#include <grp.h>
*/
import "C"

func main() {}

// Functions of the type: _nss_aad_setDBent
//
// These functions prepare the service for following operations. For a simple file based lookup this
// means files could be opened, for other services these functions are simply a noop. One special
// case for these functions is that they take an additional argument for some databases (i.e., the
// interface is int setDBent (int)). Host Names, which describes the sethostent function. The return
// value should be NSS_STATUS_SUCCESS or according to the table in "The Interface of the Function
// in NSS Modules" in case of an error.

//export _nss_aad_setpwent
func _nss_aad_setpwent() C.nss_status {
	return C.NSS_STATUS_SUCCESS
}

// Functions of the type: _nss_aad_endDBent
//
// These functions simply close all files which are still open or removes buffer caches. If there
// are no files or buffers to remove this is again a simple noop. There normally is no return value
// other than NSS_STATUS_SUCCESS.

//export _nss_aad_endpwent
func _nss_aad_endpwent() C.nss_status {
	return C.NSS_STATUS_SUCCESS
}

// Functions of the type: _nss_aad_getDBent_r
//
// These functions will be called several times in a row to retrieve one entry after the other so
// they must keep some kind of state. But this also means the functions are not really reentrant.
// They are reentrant only in that simultaneous calls to the function will not try to write the
// retrieved data in the same place (as it would be the case for the non-reentrant functions);
// instead, it writes to the structure pointed to by the result parameter. But the calls share a
// common state and in the case of a file access this means they return neighboring entries in the
// file. The buffer of length buflen pointed to by buffer can be used for storing some additional
// data for the result. It is not guaranteed that the same buffer will be passed for the next call
// of the function. Therefore one must not misuse this buffer to save some state information from
// one call to another. Before the function returns with a failure code, the implementation should
// store the value of the local errno variable in the variable pointed to be errnop. This is
// important to guarantee the module working in statically linked programs. The stored value must
// not be zero. As explained above the function could also have an additional last argument. This
// depends on the database used; it happens only for host and networks. The function shall return
// NSS_STATUS_SUCCESS as long as there are more entries. When the last entry was read it should
// return NSS_STATUS_NOTFOUND. When the buffer given as an argument is too small for the data to be
// returned NSS_STATUS_TRYAGAIN should be returned. When the service was not formerly initialized by
// a call to _nss_DATABASE_setdbent all return values allowed for this function can also be returned
// here.

//export _nss_aad_getpwent_r
func _nss_aad_getpwent_r(result *C.passwd, buffer *C.char, buflen C.size_t, errnop C.int) C.nss_status {
	return C.NSS_STATUS_NOTFOUND
}

// Functions of the type: _nss_aad_getdbbyXX_r
//
// These functions shall return the entry from the database which is addressed by the PARAMS. The
// type and number of these arguments vary. It must be individually determined by looking to the
// user-level interface functions. All arguments given to the non-reentrant version are here
// described by PARAMS. The result must be stored in the structure pointed to by result. If there
// are additional data to return (say strings, where the result structure only contains pointers)
// the function must use the buffer of length buflen. There must not be any references to
// non-constant global data. The implementation of this function should honor the stayopen flag set
// by the setDBent function whenever this makes sense. Before the function returns, the
// implementation should store the value of the local errno variable in the variable pointed to by
// errnop. This is important to guarantee the module works in statically linked programs. Again,
// this function takes an additional last argument for the host and networks database. The return
// value should as always follow the rules given in "The Interface of the Function in NSS Modules".

//export _nss_aad_getpwbyuid_r
func _nss_aad_getpwbyuid_r(uid *C.uid_t, result *C.passwd, buffer *C.char, buflen C.size_t, errnop C.int) C.nss_status {
	return C.NSS_STATUS_SUCCESS
}

//export _nss_aad_getpwbynam_r
func _nss_aad_getpwbynam_r(name *C.char, result *C.passwd, buffer *C.char, buflen C.size_t, errnop C.int) C.nss_status {
	// See documentation for _nss_aad_getpwbyuid_r
	return C.NSS_STATUS_SUCCESS
}
