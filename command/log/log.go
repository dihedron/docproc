package log

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
)

type Level int

// silent_lvl=0
// crt_lvl=1
// err_lvl=2
// wrn_lvl=3
// ntf_lvl=4
// inf_lvl=5
// dbg_lvl=6

const (
	AlwaysLog Level = iota
	CriticalLog
	ErrorLog
	WarningLog
	NotificationLog
	InformationalLog
	DebugLog
)

const DefaultLogLevel Level = InformationalLog

// Log is the command that logs at the given level.
type Log struct {
	Command
	// Level is the logging level.
	Level Level `short:"l" long:"level" description:"The log level" optional:"yes" default:"1" env:"_BENV_LOG_LEVEL"`

	Format string `short:"f" long:"format" description:"The log level format" choice:"long" choice:"medium" choice:"short" optional:"yes" default:"medium"`
}

func (cmd *Log) Execute(args []string) error {
	EnvLogLevel := DefaultLogLevel
	l := os.Getenv("_BENV_LOG_LEVEL")
	if l != "" {
		if c, err := strconv.Atoi(l); err != nil {
			return fmt.Errorf("invalid log level '%s': %w", l, err)
		} else {
			EnvLogLevel = Level(c)
		}
	}

	// if [ $_BENV_LOG_LEVEL -ge $verb_lvl ]; then
	// 	datestring=`date +"%Y-%m-%d %H:%M:%S"` >&2 echo -e "$datestring - $@"
	// fi
	if cmd.Level <= EnvLogLevel {
		now := time.Now().Format("2006-01-02 15:04:05")
		// colblk='\033[0;30m' # Black - Regular
		// colred='\033[0;31m' # Red
		// colgrn='\033[0;32m' # Green
		// colylw='\033[0;33m' # Yellow
		// colpur='\033[0;35m' # Purple
		// colrst='\033[0m'    # Text Reset
		// function esilent () { verb_lvl=$silent_lvl elog "${FUNCNAME[1]}: $@" ;}
		// function enotify () { verb_lvl=$ntf_lvl elog "${FUNCNAME[1]}: $@" ;}
		// function eok ()    { verb_lvl=$ntf_lvl elog "${colgrn}[SUCCESS]${colrst} --- ${FUNCNAME[1]}: $@" ;}
		// function ewarn ()  { verb_lvl=$wrn_lvl elog "${colylw}[WARNING]${colrst} --- ${FUNCNAME[1]}: $@" ;}
		// function einfo ()  { verb_lvl=$inf_lvl elog "${colwht}[INFO]${colrst} --- ${FUNCNAME[1]}: $@" ;}
		// function edebug () { verb_lvl=$dbg_lvl elog "[DEBUG] --- ${FUNCNAME[1]}: $@" ;}
		// function eerror () { verb_lvl=$err_lvl elog "${colred}[ERROR]${colrst} --- ${FUNCNAME[1]}: $@" ;}
		// function ecrit ()  { verb_lvl=$crt_lvl elog "${colpur}[CRITICAL]${colrst} --- ${FUNCNAME[1]}: $@" ;}
		isTTY := false
		if isatty.IsTerminal(os.Stdout.Fd()) {
			isTTY = true
		}
		caller := ""
		if cmd.Caller != "" {
			caller = fmt.Sprintf(" --- %s", cmd.Caller)
		}
		severity := ""
		switch Level(cmd.Level) {
		case AlwaysLog:
			switch cmd.Format {
			case "long":
				severity = "ALWAYS"
			case "medium":
				severity = "ALW"
			case "short":
				severity = "A"
			}
			if isTTY {
				severity = fmt.Sprintf("%s - [%s]%s:", now, color.HiBlueString(severity), caller)
			} else {
				severity = fmt.Sprintf("%s - [%s]%s:", now, severity, caller)
			}
		case CriticalLog:
			switch cmd.Format {
			case "long":
				severity = "CRITICAL"
			case "medium":
				severity = "CRT"
			case "short":
				severity = "C"
			}
			if isTTY {
				severity = fmt.Sprintf("%s - [%s]%s:", now, color.HiMagentaString(severity), caller)
			} else {
				severity = fmt.Sprintf("%s - [%s]%s:", now, severity, caller)
			}
		case ErrorLog:
			switch cmd.Format {
			case "long":
				severity = "ERROR"
			case "medium":
				severity = "ERR"
			case "short":
				severity = "E"
			}
			if isTTY {
				severity = fmt.Sprintf("%s - [%s]%s:", now, color.HiRedString(severity), caller)
			} else {
				severity = fmt.Sprintf("%s - [%s]%s:", now, severity, caller)
			}
		case WarningLog:
			switch cmd.Format {
			case "long":
				severity = "WARNING"
			case "medium":
				severity = "WRN"
			case "short":
				severity = "W"
			}
			if isTTY {
				severity = fmt.Sprintf("%s - [%s]%s:", now, color.HiYellowString(severity), caller)
			} else {
				severity = fmt.Sprintf("%s - [%s]%s:", now, severity, caller)
			}
		case NotificationLog:
			switch cmd.Format {
			case "long":
				severity = "SUCCESS"
			case "medium":
				severity = "SUC"
			case "short":
				severity = "S"
			}
			if isTTY {
				severity = fmt.Sprintf("%s - [%s]%s:", now, color.HiGreenString(severity), caller)
			} else {
				severity = fmt.Sprintf("%s - [%s]%s:", now, severity, caller)
			}
		case InformationalLog:
			switch cmd.Format {
			case "long":
				severity = "INFORMATIONAL"
			case "medium":
				severity = "INF"
			case "short":
				severity = "I"
			}
			if isTTY {
				severity = fmt.Sprintf("%s - [%s]%s:", now, color.HiWhiteString(severity), caller)
			} else {

			}
		case DebugLog:
			switch cmd.Format {
			case "long":
				severity = "DEBUG"
			case "medium":
				severity = "DBG"
			case "short":
				severity = "D"
			}
			severity = fmt.Sprintf("%s - [%s]%s:", now, severity, caller)
		}
		fmt.Printf("%s %v\n", severity, strings.Join(args, " "))
	}
	return nil
}
