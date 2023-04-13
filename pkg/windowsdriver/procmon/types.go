//go:build ignore
// +build ignore

package procmon

/*
//! These includes are needed to use constants defined in the ddprocmonapi
#include <WinDef.h>
#include <WinIoCtl.h>

//! Defines the objects used to communicate with the driver as well as its control codes
#include "../include/procmonapi.h"
#include <stdlib.h>
*/
import "C"

const Signature = C.DD_PROCMONDRIVER_SIGNATURE

const (
	ProcmonNotifyStop  = C.DD_NOTIFY_STOP
	ProcmonNotifyStart = C.DD_NOTIFY_START
)

type DDProcessNotifyType C.enum__dd_notify_type
type DDProcessNotification C.struct__dd_process_notification
