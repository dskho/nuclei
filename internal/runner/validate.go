package runner

import (
	"errors"
	"net/url"

	"github.com/projectdiscovery/gologger"
)

// validateOptions validates the configuration options passed
func (options *Options) validateOptions() error {
	// Both verbose and silent flags were used
	if options.Verbose && options.Silent {
		return errors.New("both verbose and silent mode specified")
	}

	// Check if a list of templates was provided and it exists
	if len(options.Templates) == 0 && !options.UpdateTemplates {
		return errors.New("no template/templates provided")
	}

	if options.Targets == "" && !options.Stdin && options.Target == "" && !options.UpdateTemplates {
		return errors.New("no target input provided")
	}

	// Validate proxy options if provided
	if options.ProxyURL != "" && !isValidProxyURL(options.ProxyURL) {
		return errors.New("invalid http proxy format (It should be http://username:password@host:port)")
	}

	if options.ProxySocksURL != "" && !isValidProxyURL(options.ProxySocksURL) {
		return errors.New("invalid socks proxy format (It should be socks5://username:password@host:port)")
	}

	return nil
}

func isValidProxyURL(proxyURL string) bool {
	_, err := url.Parse(proxyURL)

	return err == nil
}

// configureOutput configures the output on the screen
func (options *Options) configureOutput() {
	// If the user desires verbose output, show verbose output
	if options.Verbose {
		gologger.MaxLevel = gologger.Verbose
	}

	if options.NoColor {
		gologger.UseColors = false
	}

	if options.Silent {
		gologger.MaxLevel = gologger.Silent
	}
}
