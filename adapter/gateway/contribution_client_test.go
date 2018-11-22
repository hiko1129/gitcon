package gateway_test

import (
	"github.com/hiko1129/gitcon/infrastructure"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hiko1129/gitcon/adapter/gateway"
)

// TODO remove httpmock

var _ = Describe("ContributionClient", func() {
	var client *gateway.ContirubutionClient
	BeforeEach(func() {
		http := infrastructure.NewHTTP()
		client = gateway.NewContirubutionClient(http)
	})

	Describe("#FetchContributions", func() {
		html := `<svg width="563" height="88" class="js-calendar-graph-svg">
  <g transform="translate(16, 20)">
      <g transform="translate(0, 0)">
          <rect class="day" width="8" height="8" x="11" y="0" fill="#7bc96f" data-count="13" data-date="2017-11-19"></rect>
          <rect class="day" width="8" height="8" x="11" y="10" fill="#7bc96f" data-count="13" data-date="2017-11-20"></rect>
          <rect class="day" width="8" height="8" x="11" y="20" fill="#7bc96f" data-count="10" data-date="2017-11-21"></rect>
          <rect class="day" width="8" height="8" x="11" y="30" fill="#c6e48b" data-count="5" data-date="2017-11-22"></rect>
          <rect class="day" width="8" height="8" x="11" y="40" fill="#ebedf0" data-count="0" data-date="2017-11-23"></rect>
          <rect class="day" width="8" height="8" x="11" y="50" fill="#ebedf0" data-count="0" data-date="2017-11-24"></rect>
          <rect class="day" width="8" height="8" x="11" y="60" fill="#ebedf0" data-count="0" data-date="2017-11-25"></rect>
      </g>
      <g transform="translate(11, 0)">
          <rect class="day" width="8" height="8" x="10" y="0" fill="#ebedf0" data-count="0" data-date="2017-11-26"></rect>
          <rect class="day" width="8" height="8" x="10" y="10" fill="#c6e48b" data-count="3" data-date="2017-11-27"></rect>
          <rect class="day" width="8" height="8" x="10" y="20" fill="#7bc96f" data-count="10" data-date="2017-11-28"></rect>
          <rect class="day" width="8" height="8" x="10" y="30" fill="#7bc96f" data-count="11" data-date="2017-11-29"></rect>
          <rect class="day" width="8" height="8" x="10" y="40" fill="#c6e48b" data-count="6" data-date="2017-11-30"></rect>
          <rect class="day" width="8" height="8" x="10" y="50" fill="#c6e48b" data-count="7" data-date="2017-12-01"></rect>
          <rect class="day" width="8" height="8" x="10" y="60" fill="#c6e48b" data-count="2" data-date="2017-12-02"></rect>
      </g>
      <g transform="translate(22, 0)">
          <rect class="day" width="8" height="8" x="9" y="0" fill="#c6e48b" data-count="1" data-date="2017-12-03"></rect>
          <rect class="day" width="8" height="8" x="9" y="10" fill="#c6e48b" data-count="9" data-date="2017-12-04"></rect>
          <rect class="day" width="8" height="8" x="9" y="20" fill="#c6e48b" data-count="6" data-date="2017-12-05"></rect>
          <rect class="day" width="8" height="8" x="9" y="30" fill="#c6e48b" data-count="9" data-date="2017-12-06"></rect>
          <rect class="day" width="8" height="8" x="9" y="40" fill="#c6e48b" data-count="9" data-date="2017-12-07"></rect>
          <rect class="day" width="8" height="8" x="9" y="50" fill="#7bc96f" data-count="12" data-date="2017-12-08"></rect>
          <rect class="day" width="8" height="8" x="9" y="60" fill="#c6e48b" data-count="2" data-date="2017-12-09"></rect>
      </g>
      <g transform="translate(33, 0)">
          <rect class="day" width="8" height="8" x="8" y="0" fill="#7bc96f" data-count="15" data-date="2017-12-10"></rect>
          <rect class="day" width="8" height="8" x="8" y="10" fill="#7bc96f" data-count="10" data-date="2017-12-11"></rect>
          <rect class="day" width="8" height="8" x="8" y="20" fill="#7bc96f" data-count="12" data-date="2017-12-12"></rect>
          <rect class="day" width="8" height="8" x="8" y="30" fill="#7bc96f" data-count="12" data-date="2017-12-13"></rect>
          <rect class="day" width="8" height="8" x="8" y="40" fill="#7bc96f" data-count="10" data-date="2017-12-14"></rect>
          <rect class="day" width="8" height="8" x="8" y="50" fill="#7bc96f" data-count="12" data-date="2017-12-15"></rect>
          <rect class="day" width="8" height="8" x="8" y="60" fill="#c6e48b" data-count="2" data-date="2017-12-16"></rect>
      </g>
      <g transform="translate(44, 0)">
          <rect class="day" width="8" height="8" x="7" y="0" fill="#ebedf0" data-count="0" data-date="2017-12-17"></rect>
          <rect class="day" width="8" height="8" x="7" y="10" fill="#c6e48b" data-count="3" data-date="2017-12-18"></rect>
          <rect class="day" width="8" height="8" x="7" y="20" fill="#c6e48b" data-count="8" data-date="2017-12-19"></rect>
          <rect class="day" width="8" height="8" x="7" y="30" fill="#7bc96f" data-count="17" data-date="2017-12-20"></rect>
          <rect class="day" width="8" height="8" x="7" y="40" fill="#239a3b" data-count="22" data-date="2017-12-21"></rect>
          <rect class="day" width="8" height="8" x="7" y="50" fill="#7bc96f" data-count="15" data-date="2017-12-22"></rect>
          <rect class="day" width="8" height="8" x="7" y="60" fill="#ebedf0" data-count="0" data-date="2017-12-23"></rect>
      </g>
      <g transform="translate(55, 0)">
          <rect class="day" width="8" height="8" x="6" y="0" fill="#ebedf0" data-count="0" data-date="2017-12-24"></rect>
          <rect class="day" width="8" height="8" x="6" y="10" fill="#239a3b" data-count="20" data-date="2017-12-25"></rect>
          <rect class="day" width="8" height="8" x="6" y="20" fill="#7bc96f" data-count="19" data-date="2017-12-26"></rect>
          <rect class="day" width="8" height="8" x="6" y="30" fill="#c6e48b" data-count="9" data-date="2017-12-27"></rect>
          <rect class="day" width="8" height="8" x="6" y="40" fill="#c6e48b" data-count="7" data-date="2017-12-28"></rect>
          <rect class="day" width="8" height="8" x="6" y="50" fill="#ebedf0" data-count="0" data-date="2017-12-29"></rect>
          <rect class="day" width="8" height="8" x="6" y="60" fill="#ebedf0" data-count="0" data-date="2017-12-30"></rect>
      </g>
      <g transform="translate(66, 0)">
          <rect class="day" width="8" height="8" x="5" y="0" fill="#ebedf0" data-count="0" data-date="2017-12-31"></rect>
          <rect class="day" width="8" height="8" x="5" y="10" fill="#ebedf0" data-count="0" data-date="2018-01-01"></rect>
          <rect class="day" width="8" height="8" x="5" y="20" fill="#ebedf0" data-count="0" data-date="2018-01-02"></rect>
          <rect class="day" width="8" height="8" x="5" y="30" fill="#ebedf0" data-count="0" data-date="2018-01-03"></rect>
          <rect class="day" width="8" height="8" x="5" y="40" fill="#ebedf0" data-count="0" data-date="2018-01-04"></rect>
          <rect class="day" width="8" height="8" x="5" y="50" fill="#7bc96f" data-count="14" data-date="2018-01-05"></rect>
          <rect class="day" width="8" height="8" x="5" y="60" fill="#ebedf0" data-count="0" data-date="2018-01-06"></rect>
      </g>
      <g transform="translate(77, 0)">
          <rect class="day" width="8" height="8" x="4" y="0" fill="#ebedf0" data-count="0" data-date="2018-01-07"></rect>
          <rect class="day" width="8" height="8" x="4" y="10" fill="#ebedf0" data-count="0" data-date="2018-01-08"></rect>
          <rect class="day" width="8" height="8" x="4" y="20" fill="#c6e48b" data-count="4" data-date="2018-01-09"></rect>
          <rect class="day" width="8" height="8" x="4" y="30" fill="#c6e48b" data-count="7" data-date="2018-01-10"></rect>
          <rect class="day" width="8" height="8" x="4" y="40" fill="#c6e48b" data-count="9" data-date="2018-01-11"></rect>
          <rect class="day" width="8" height="8" x="4" y="50" fill="#7bc96f" data-count="19" data-date="2018-01-12"></rect>
          <rect class="day" width="8" height="8" x="4" y="60" fill="#ebedf0" data-count="0" data-date="2018-01-13"></rect>
      </g>
      <g transform="translate(88, 0)">
          <rect class="day" width="8" height="8" x="3" y="0" fill="#ebedf0" data-count="0" data-date="2018-01-14"></rect>
          <rect class="day" width="8" height="8" x="3" y="10" fill="#239a3b" data-count="26" data-date="2018-01-15"></rect>
          <rect class="day" width="8" height="8" x="3" y="20" fill="#7bc96f" data-count="15" data-date="2018-01-16"></rect>
          <rect class="day" width="8" height="8" x="3" y="30" fill="#ebedf0" data-count="0" data-date="2018-01-17"></rect>
          <rect class="day" width="8" height="8" x="3" y="40" fill="#c6e48b" data-count="5" data-date="2018-01-18"></rect>
          <rect class="day" width="8" height="8" x="3" y="50" fill="#7bc96f" data-count="16" data-date="2018-01-19"></rect>
          <rect class="day" width="8" height="8" x="3" y="60" fill="#c6e48b" data-count="3" data-date="2018-01-20"></rect>
      </g>
      <g transform="translate(99, 0)">
          <rect class="day" width="8" height="8" x="2" y="0" fill="#c6e48b" data-count="5" data-date="2018-01-21"></rect>
          <rect class="day" width="8" height="8" x="2" y="10" fill="#7bc96f" data-count="12" data-date="2018-01-22"></rect>
          <rect class="day" width="8" height="8" x="2" y="20" fill="#7bc96f" data-count="18" data-date="2018-01-23"></rect>
          <rect class="day" width="8" height="8" x="2" y="30" fill="#c6e48b" data-count="1" data-date="2018-01-24"></rect>
          <rect class="day" width="8" height="8" x="2" y="40" fill="#7bc96f" data-count="12" data-date="2018-01-25"></rect>
          <rect class="day" width="8" height="8" x="2" y="50" fill="#7bc96f" data-count="13" data-date="2018-01-26"></rect>
          <rect class="day" width="8" height="8" x="2" y="60" fill="#ebedf0" data-count="0" data-date="2018-01-27"></rect>
      </g>
      <g transform="translate(110, 0)">
          <rect class="day" width="8" height="8" x="1" y="0" fill="#c6e48b" data-count="3" data-date="2018-01-28"></rect>
          <rect class="day" width="8" height="8" x="1" y="10" fill="#7bc96f" data-count="14" data-date="2018-01-29"></rect>
          <rect class="day" width="8" height="8" x="1" y="20" fill="#7bc96f" data-count="15" data-date="2018-01-30"></rect>
          <rect class="day" width="8" height="8" x="1" y="30" fill="#7bc96f" data-count="15" data-date="2018-01-31"></rect>
          <rect class="day" width="8" height="8" x="1" y="40" fill="#c6e48b" data-count="9" data-date="2018-02-01"></rect>
          <rect class="day" width="8" height="8" x="1" y="50" fill="#c6e48b" data-count="6" data-date="2018-02-02"></rect>
          <rect class="day" width="8" height="8" x="1" y="60" fill="#ebedf0" data-count="0" data-date="2018-02-03"></rect>
      </g>
      <g transform="translate(121, 0)">
          <rect class="day" width="8" height="8" x="0" y="0" fill="#ebedf0" data-count="0" data-date="2018-02-04"></rect>
          <rect class="day" width="8" height="8" x="0" y="10" fill="#239a3b" data-count="21" data-date="2018-02-05"></rect>
          <rect class="day" width="8" height="8" x="0" y="20" fill="#c6e48b" data-count="8" data-date="2018-02-06"></rect>
          <rect class="day" width="8" height="8" x="0" y="30" fill="#7bc96f" data-count="11" data-date="2018-02-07"></rect>
          <rect class="day" width="8" height="8" x="0" y="40" fill="#239a3b" data-count="21" data-date="2018-02-08"></rect>
          <rect class="day" width="8" height="8" x="0" y="50" fill="#c6e48b" data-count="6" data-date="2018-02-09"></rect>
          <rect class="day" width="8" height="8" x="0" y="60" fill="#ebedf0" data-count="0" data-date="2018-02-10"></rect>
      </g>
      <g transform="translate(132, 0)">
          <rect class="day" width="8" height="8" x="-1" y="0" fill="#ebedf0" data-count="0" data-date="2018-02-11"></rect>
          <rect class="day" width="8" height="8" x="-1" y="10" fill="#ebedf0" data-count="0" data-date="2018-02-12"></rect>
          <rect class="day" width="8" height="8" x="-1" y="20" fill="#c6e48b" data-count="8" data-date="2018-02-13"></rect>
          <rect class="day" width="8" height="8" x="-1" y="30" fill="#c6e48b" data-count="4" data-date="2018-02-14"></rect>
          <rect class="day" width="8" height="8" x="-1" y="40" fill="#7bc96f" data-count="11" data-date="2018-02-15"></rect>
          <rect class="day" width="8" height="8" x="-1" y="50" fill="#7bc96f" data-count="17" data-date="2018-02-16"></rect>
          <rect class="day" width="8" height="8" x="-1" y="60" fill="#c6e48b" data-count="3" data-date="2018-02-17"></rect>
      </g>
      <g transform="translate(143, 0)">
          <rect class="day" width="8" height="8" x="-2" y="0" fill="#ebedf0" data-count="0" data-date="2018-02-18"></rect>
          <rect class="day" width="8" height="8" x="-2" y="10" fill="#196127" data-count="33" data-date="2018-02-19"></rect>
          <rect class="day" width="8" height="8" x="-2" y="20" fill="#c6e48b" data-count="6" data-date="2018-02-20"></rect>
          <rect class="day" width="8" height="8" x="-2" y="30" fill="#7bc96f" data-count="18" data-date="2018-02-21"></rect>
          <rect class="day" width="8" height="8" x="-2" y="40" fill="#239a3b" data-count="27" data-date="2018-02-22"></rect>
          <rect class="day" width="8" height="8" x="-2" y="50" fill="#239a3b" data-count="22" data-date="2018-02-23"></rect>
          <rect class="day" width="8" height="8" x="-2" y="60" fill="#c6e48b" data-count="1" data-date="2018-02-24"></rect>
      </g>
      <g transform="translate(154, 0)">
          <rect class="day" width="8" height="8" x="-3" y="0" fill="#ebedf0" data-count="0" data-date="2018-02-25"></rect>
          <rect class="day" width="8" height="8" x="-3" y="10" fill="#c6e48b" data-count="3" data-date="2018-02-26"></rect>
          <rect class="day" width="8" height="8" x="-3" y="20" fill="#c6e48b" data-count="6" data-date="2018-02-27"></rect>
          <rect class="day" width="8" height="8" x="-3" y="30" fill="#ebedf0" data-count="0" data-date="2018-02-28"></rect>
          <rect class="day" width="8" height="8" x="-3" y="40" fill="#c6e48b" data-count="7" data-date="2018-03-01"></rect>
          <rect class="day" width="8" height="8" x="-3" y="50" fill="#c6e48b" data-count="7" data-date="2018-03-02"></rect>
          <rect class="day" width="8" height="8" x="-3" y="60" fill="#ebedf0" data-count="0" data-date="2018-03-03"></rect>
      </g>
      <g transform="translate(165, 0)">
          <rect class="day" width="8" height="8" x="-4" y="0" fill="#ebedf0" data-count="0" data-date="2018-03-04"></rect>
          <rect class="day" width="8" height="8" x="-4" y="10" fill="#7bc96f" data-count="10" data-date="2018-03-05"></rect>
          <rect class="day" width="8" height="8" x="-4" y="20" fill="#239a3b" data-count="24" data-date="2018-03-06"></rect>
          <rect class="day" width="8" height="8" x="-4" y="30" fill="#196127" data-count="39" data-date="2018-03-07"></rect>
          <rect class="day" width="8" height="8" x="-4" y="40" fill="#239a3b" data-count="21" data-date="2018-03-08"></rect>
          <rect class="day" width="8" height="8" x="-4" y="50" fill="#239a3b" data-count="20" data-date="2018-03-09"></rect>
          <rect class="day" width="8" height="8" x="-4" y="60" fill="#ebedf0" data-count="0" data-date="2018-03-10"></rect>
      </g>
      <g transform="translate(176, 0)">
          <rect class="day" width="8" height="8" x="-5" y="0" fill="#ebedf0" data-count="0" data-date="2018-03-11"></rect>
          <rect class="day" width="8" height="8" x="-5" y="10" fill="#ebedf0" data-count="0" data-date="2018-03-12"></rect>
          <rect class="day" width="8" height="8" x="-5" y="20" fill="#ebedf0" data-count="0" data-date="2018-03-13"></rect>
          <rect class="day" width="8" height="8" x="-5" y="30" fill="#ebedf0" data-count="0" data-date="2018-03-14"></rect>
          <rect class="day" width="8" height="8" x="-5" y="40" fill="#7bc96f" data-count="10" data-date="2018-03-15"></rect>
          <rect class="day" width="8" height="8" x="-5" y="50" fill="#196127" data-count="32" data-date="2018-03-16"></rect>
          <rect class="day" width="8" height="8" x="-5" y="60" fill="#ebedf0" data-count="0" data-date="2018-03-17"></rect>
      </g>
      <g transform="translate(187, 0)">
          <rect class="day" width="8" height="8" x="-6" y="0" fill="#ebedf0" data-count="0" data-date="2018-03-18"></rect>
          <rect class="day" width="8" height="8" x="-6" y="10" fill="#c6e48b" data-count="8" data-date="2018-03-19"></rect>
          <rect class="day" width="8" height="8" x="-6" y="20" fill="#c6e48b" data-count="2" data-date="2018-03-20"></rect>
          <rect class="day" width="8" height="8" x="-6" y="30" fill="#ebedf0" data-count="0" data-date="2018-03-21"></rect>
          <rect class="day" width="8" height="8" x="-6" y="40" fill="#7bc96f" data-count="13" data-date="2018-03-22"></rect>
          <rect class="day" width="8" height="8" x="-6" y="50" fill="#239a3b" data-count="25" data-date="2018-03-23"></rect>
          <rect class="day" width="8" height="8" x="-6" y="60" fill="#c6e48b" data-count="1" data-date="2018-03-24"></rect>
      </g>
      <g transform="translate(198, 0)">
          <rect class="day" width="8" height="8" x="-7" y="0" fill="#ebedf0" data-count="0" data-date="2018-03-25"></rect>
          <rect class="day" width="8" height="8" x="-7" y="10" fill="#196127" data-count="34" data-date="2018-03-26"></rect>
          <rect class="day" width="8" height="8" x="-7" y="20" fill="#c6e48b" data-count="6" data-date="2018-03-27"></rect>
          <rect class="day" width="8" height="8" x="-7" y="30" fill="#7bc96f" data-count="10" data-date="2018-03-28"></rect>
          <rect class="day" width="8" height="8" x="-7" y="40" fill="#196127" data-count="31" data-date="2018-03-29"></rect>
          <rect class="day" width="8" height="8" x="-7" y="50" fill="#239a3b" data-count="27" data-date="2018-03-30"></rect>
          <rect class="day" width="8" height="8" x="-7" y="60" fill="#c6e48b" data-count="3" data-date="2018-03-31"></rect>
      </g>
      <g transform="translate(209, 0)">
          <rect class="day" width="8" height="8" x="-8" y="0" fill="#c6e48b" data-count="2" data-date="2018-04-01"></rect>
          <rect class="day" width="8" height="8" x="-8" y="10" fill="#c6e48b" data-count="7" data-date="2018-04-02"></rect>
          <rect class="day" width="8" height="8" x="-8" y="20" fill="#7bc96f" data-count="10" data-date="2018-04-03"></rect>
          <rect class="day" width="8" height="8" x="-8" y="30" fill="#7bc96f" data-count="15" data-date="2018-04-04"></rect>
          <rect class="day" width="8" height="8" x="-8" y="40" fill="#c6e48b" data-count="9" data-date="2018-04-05"></rect>
          <rect class="day" width="8" height="8" x="-8" y="50" fill="#c6e48b" data-count="9" data-date="2018-04-06"></rect>
          <rect class="day" width="8" height="8" x="-8" y="60" fill="#c6e48b" data-count="2" data-date="2018-04-07"></rect>
      </g>
      <g transform="translate(220, 0)">
          <rect class="day" width="8" height="8" x="-9" y="0" fill="#ebedf0" data-count="0" data-date="2018-04-08"></rect>
          <rect class="day" width="8" height="8" x="-9" y="10" fill="#7bc96f" data-count="16" data-date="2018-04-09"></rect>
          <rect class="day" width="8" height="8" x="-9" y="20" fill="#c6e48b" data-count="6" data-date="2018-04-10"></rect>
          <rect class="day" width="8" height="8" x="-9" y="30" fill="#c6e48b" data-count="8" data-date="2018-04-11"></rect>
          <rect class="day" width="8" height="8" x="-9" y="40" fill="#ebedf0" data-count="0" data-date="2018-04-12"></rect>
          <rect class="day" width="8" height="8" x="-9" y="50" fill="#239a3b" data-count="28" data-date="2018-04-13"></rect>
          <rect class="day" width="8" height="8" x="-9" y="60" fill="#ebedf0" data-count="0" data-date="2018-04-14"></rect>
      </g>
      <g transform="translate(231, 0)">
          <rect class="day" width="8" height="8" x="-10" y="0" fill="#ebedf0" data-count="0" data-date="2018-04-15"></rect>
          <rect class="day" width="8" height="8" x="-10" y="10" fill="#239a3b" data-count="20" data-date="2018-04-16"></rect>
          <rect class="day" width="8" height="8" x="-10" y="20" fill="#239a3b" data-count="24" data-date="2018-04-17"></rect>
          <rect class="day" width="8" height="8" x="-10" y="30" fill="#7bc96f" data-count="13" data-date="2018-04-18"></rect>
          <rect class="day" width="8" height="8" x="-10" y="40" fill="#7bc96f" data-count="10" data-date="2018-04-19"></rect>
          <rect class="day" width="8" height="8" x="-10" y="50" fill="#7bc96f" data-count="19" data-date="2018-04-20"></rect>
          <rect class="day" width="8" height="8" x="-10" y="60" fill="#ebedf0" data-count="0" data-date="2018-04-21"></rect>
      </g>
      <g transform="translate(242, 0)">
          <rect class="day" width="8" height="8" x="-11" y="0" fill="#ebedf0" data-count="0" data-date="2018-04-22"></rect>
          <rect class="day" width="8" height="8" x="-11" y="10" fill="#239a3b" data-count="22" data-date="2018-04-23"></rect>
          <rect class="day" width="8" height="8" x="-11" y="20" fill="#7bc96f" data-count="13" data-date="2018-04-24"></rect>
          <rect class="day" width="8" height="8" x="-11" y="30" fill="#7bc96f" data-count="15" data-date="2018-04-25"></rect>
          <rect class="day" width="8" height="8" x="-11" y="40" fill="#239a3b" data-count="23" data-date="2018-04-26"></rect>
          <rect class="day" width="8" height="8" x="-11" y="50" fill="#196127" data-count="37" data-date="2018-04-27"></rect>
          <rect class="day" width="8" height="8" x="-11" y="60" fill="#ebedf0" data-count="0" data-date="2018-04-28"></rect>
      </g>
      <g transform="translate(253, 0)">
          <rect class="day" width="8" height="8" x="-12" y="0" fill="#ebedf0" data-count="0" data-date="2018-04-29"></rect>
          <rect class="day" width="8" height="8" x="-12" y="10" fill="#ebedf0" data-count="0" data-date="2018-04-30"></rect>
          <rect class="day" width="8" height="8" x="-12" y="20" fill="#ebedf0" data-count="0" data-date="2018-05-01"></rect>
          <rect class="day" width="8" height="8" x="-12" y="30" fill="#ebedf0" data-count="0" data-date="2018-05-02"></rect>
          <rect class="day" width="8" height="8" x="-12" y="40" fill="#ebedf0" data-count="0" data-date="2018-05-03"></rect>
          <rect class="day" width="8" height="8" x="-12" y="50" fill="#c6e48b" data-count="8" data-date="2018-05-04"></rect>
          <rect class="day" width="8" height="8" x="-12" y="60" fill="#c6e48b" data-count="1" data-date="2018-05-05"></rect>
      </g>
      <g transform="translate(264, 0)">
          <rect class="day" width="8" height="8" x="-13" y="0" fill="#ebedf0" data-count="0" data-date="2018-05-06"></rect>
          <rect class="day" width="8" height="8" x="-13" y="10" fill="#ebedf0" data-count="0" data-date="2018-05-07"></rect>
          <rect class="day" width="8" height="8" x="-13" y="20" fill="#7bc96f" data-count="17" data-date="2018-05-08"></rect>
          <rect class="day" width="8" height="8" x="-13" y="30" fill="#c6e48b" data-count="6" data-date="2018-05-09"></rect>
          <rect class="day" width="8" height="8" x="-13" y="40" fill="#196127" data-count="31" data-date="2018-05-10"></rect>
          <rect class="day" width="8" height="8" x="-13" y="50" fill="#c6e48b" data-count="7" data-date="2018-05-11"></rect>
          <rect class="day" width="8" height="8" x="-13" y="60" fill="#ebedf0" data-count="0" data-date="2018-05-12"></rect>
      </g>
      <g transform="translate(275, 0)">
          <rect class="day" width="8" height="8" x="-14" y="0" fill="#ebedf0" data-count="0" data-date="2018-05-13"></rect>
          <rect class="day" width="8" height="8" x="-14" y="10" fill="#196127" data-count="43" data-date="2018-05-14"></rect>
          <rect class="day" width="8" height="8" x="-14" y="20" fill="#c6e48b" data-count="9" data-date="2018-05-15"></rect>
          <rect class="day" width="8" height="8" x="-14" y="30" fill="#c6e48b" data-count="9" data-date="2018-05-16"></rect>
          <rect class="day" width="8" height="8" x="-14" y="40" fill="#7bc96f" data-count="10" data-date="2018-05-17"></rect>
          <rect class="day" width="8" height="8" x="-14" y="50" fill="#196127" data-count="32" data-date="2018-05-18"></rect>
          <rect class="day" width="8" height="8" x="-14" y="60" fill="#ebedf0" data-count="0" data-date="2018-05-19"></rect>
      </g>
      <g transform="translate(286, 0)">
          <rect class="day" width="8" height="8" x="-15" y="0" fill="#ebedf0" data-count="0" data-date="2018-05-20"></rect>
          <rect class="day" width="8" height="8" x="-15" y="10" fill="#c6e48b" data-count="2" data-date="2018-05-21"></rect>
          <rect class="day" width="8" height="8" x="-15" y="20" fill="#c6e48b" data-count="4" data-date="2018-05-22"></rect>
          <rect class="day" width="8" height="8" x="-15" y="30" fill="#ebedf0" data-count="0" data-date="2018-05-23"></rect>
          <rect class="day" width="8" height="8" x="-15" y="40" fill="#c6e48b" data-count="2" data-date="2018-05-24"></rect>
          <rect class="day" width="8" height="8" x="-15" y="50" fill="#c6e48b" data-count="3" data-date="2018-05-25"></rect>
          <rect class="day" width="8" height="8" x="-15" y="60" fill="#ebedf0" data-count="0" data-date="2018-05-26"></rect>
      </g>
      <g transform="translate(297, 0)">
          <rect class="day" width="8" height="8" x="-16" y="0" fill="#ebedf0" data-count="0" data-date="2018-05-27"></rect>
          <rect class="day" width="8" height="8" x="-16" y="10" fill="#c6e48b" data-count="1" data-date="2018-05-28"></rect>
          <rect class="day" width="8" height="8" x="-16" y="20" fill="#c6e48b" data-count="4" data-date="2018-05-29"></rect>
          <rect class="day" width="8" height="8" x="-16" y="30" fill="#c6e48b" data-count="2" data-date="2018-05-30"></rect>
          <rect class="day" width="8" height="8" x="-16" y="40" fill="#ebedf0" data-count="0" data-date="2018-05-31"></rect>
          <rect class="day" width="8" height="8" x="-16" y="50" fill="#ebedf0" data-count="0" data-date="2018-06-01"></rect>
          <rect class="day" width="8" height="8" x="-16" y="60" fill="#ebedf0" data-count="0" data-date="2018-06-02"></rect>
      </g>
      <g transform="translate(308, 0)">
          <rect class="day" width="8" height="8" x="-17" y="0" fill="#ebedf0" data-count="0" data-date="2018-06-03"></rect>
          <rect class="day" width="8" height="8" x="-17" y="10" fill="#ebedf0" data-count="0" data-date="2018-06-04"></rect>
          <rect class="day" width="8" height="8" x="-17" y="20" fill="#c6e48b" data-count="1" data-date="2018-06-05"></rect>
          <rect class="day" width="8" height="8" x="-17" y="30" fill="#ebedf0" data-count="0" data-date="2018-06-06"></rect>
          <rect class="day" width="8" height="8" x="-17" y="40" fill="#c6e48b" data-count="2" data-date="2018-06-07"></rect>
          <rect class="day" width="8" height="8" x="-17" y="50" fill="#c6e48b" data-count="1" data-date="2018-06-08"></rect>
          <rect class="day" width="8" height="8" x="-17" y="60" fill="#ebedf0" data-count="0" data-date="2018-06-09"></rect>
      </g>
      <g transform="translate(319, 0)">
          <rect class="day" width="8" height="8" x="-18" y="0" fill="#ebedf0" data-count="0" data-date="2018-06-10"></rect>
          <rect class="day" width="8" height="8" x="-18" y="10" fill="#c6e48b" data-count="1" data-date="2018-06-11"></rect>
          <rect class="day" width="8" height="8" x="-18" y="20" fill="#c6e48b" data-count="1" data-date="2018-06-12"></rect>
          <rect class="day" width="8" height="8" x="-18" y="30" fill="#c6e48b" data-count="2" data-date="2018-06-13"></rect>
          <rect class="day" width="8" height="8" x="-18" y="40" fill="#c6e48b" data-count="3" data-date="2018-06-14"></rect>
          <rect class="day" width="8" height="8" x="-18" y="50" fill="#c6e48b" data-count="1" data-date="2018-06-15"></rect>
          <rect class="day" width="8" height="8" x="-18" y="60" fill="#ebedf0" data-count="0" data-date="2018-06-16"></rect>
      </g>
      <g transform="translate(330, 0)">
          <rect class="day" width="8" height="8" x="-19" y="0" fill="#ebedf0" data-count="0" data-date="2018-06-17"></rect>
          <rect class="day" width="8" height="8" x="-19" y="10" fill="#c6e48b" data-count="1" data-date="2018-06-18"></rect>
          <rect class="day" width="8" height="8" x="-19" y="20" fill="#c6e48b" data-count="9" data-date="2018-06-19"></rect>
          <rect class="day" width="8" height="8" x="-19" y="30" fill="#7bc96f" data-count="12" data-date="2018-06-20"></rect>
          <rect class="day" width="8" height="8" x="-19" y="40" fill="#7bc96f" data-count="19" data-date="2018-06-21"></rect>
          <rect class="day" width="8" height="8" x="-19" y="50" fill="#7bc96f" data-count="10" data-date="2018-06-22"></rect>
          <rect class="day" width="8" height="8" x="-19" y="60" fill="#ebedf0" data-count="0" data-date="2018-06-23"></rect>
      </g>
      <g transform="translate(341, 0)">
          <rect class="day" width="8" height="8" x="-20" y="0" fill="#ebedf0" data-count="0" data-date="2018-06-24"></rect>
          <rect class="day" width="8" height="8" x="-20" y="10" fill="#ebedf0" data-count="0" data-date="2018-06-25"></rect>
          <rect class="day" width="8" height="8" x="-20" y="20" fill="#c6e48b" data-count="3" data-date="2018-06-26"></rect>
          <rect class="day" width="8" height="8" x="-20" y="30" fill="#c6e48b" data-count="9" data-date="2018-06-27"></rect>
          <rect class="day" width="8" height="8" x="-20" y="40" fill="#7bc96f" data-count="11" data-date="2018-06-28"></rect>
          <rect class="day" width="8" height="8" x="-20" y="50" fill="#239a3b" data-count="20" data-date="2018-06-29"></rect>
          <rect class="day" width="8" height="8" x="-20" y="60" fill="#ebedf0" data-count="0" data-date="2018-06-30"></rect>
      </g>
      <g transform="translate(352, 0)">
          <rect class="day" width="8" height="8" x="-21" y="0" fill="#ebedf0" data-count="0" data-date="2018-07-01"></rect>
          <rect class="day" width="8" height="8" x="-21" y="10" fill="#7bc96f" data-count="13" data-date="2018-07-02"></rect>
          <rect class="day" width="8" height="8" x="-21" y="20" fill="#c6e48b" data-count="2" data-date="2018-07-03"></rect>
          <rect class="day" width="8" height="8" x="-21" y="30" fill="#7bc96f" data-count="16" data-date="2018-07-04"></rect>
          <rect class="day" width="8" height="8" x="-21" y="40" fill="#c6e48b" data-count="6" data-date="2018-07-05"></rect>
          <rect class="day" width="8" height="8" x="-21" y="50" fill="#7bc96f" data-count="18" data-date="2018-07-06"></rect>
          <rect class="day" width="8" height="8" x="-21" y="60" fill="#ebedf0" data-count="0" data-date="2018-07-07"></rect>
      </g>
      <g transform="translate(363, 0)">
          <rect class="day" width="8" height="8" x="-22" y="0" fill="#ebedf0" data-count="0" data-date="2018-07-08"></rect>
          <rect class="day" width="8" height="8" x="-22" y="10" fill="#c6e48b" data-count="6" data-date="2018-07-09"></rect>
          <rect class="day" width="8" height="8" x="-22" y="20" fill="#c6e48b" data-count="5" data-date="2018-07-10"></rect>
          <rect class="day" width="8" height="8" x="-22" y="30" fill="#239a3b" data-count="23" data-date="2018-07-11"></rect>
          <rect class="day" width="8" height="8" x="-22" y="40" fill="#c6e48b" data-count="9" data-date="2018-07-12"></rect>
          <rect class="day" width="8" height="8" x="-22" y="50" fill="#7bc96f" data-count="10" data-date="2018-07-13"></rect>
          <rect class="day" width="8" height="8" x="-22" y="60" fill="#ebedf0" data-count="0" data-date="2018-07-14"></rect>
      </g>
      <g transform="translate(374, 0)">
          <rect class="day" width="8" height="8" x="-23" y="0" fill="#ebedf0" data-count="0" data-date="2018-07-15"></rect>
          <rect class="day" width="8" height="8" x="-23" y="10" fill="#ebedf0" data-count="0" data-date="2018-07-16"></rect>
          <rect class="day" width="8" height="8" x="-23" y="20" fill="#ebedf0" data-count="0" data-date="2018-07-17"></rect>
          <rect class="day" width="8" height="8" x="-23" y="30" fill="#c6e48b" data-count="3" data-date="2018-07-18"></rect>
          <rect class="day" width="8" height="8" x="-23" y="40" fill="#239a3b" data-count="20" data-date="2018-07-19"></rect>
          <rect class="day" width="8" height="8" x="-23" y="50" fill="#c6e48b" data-count="9" data-date="2018-07-20"></rect>
          <rect class="day" width="8" height="8" x="-23" y="60" fill="#ebedf0" data-count="0" data-date="2018-07-21"></rect>
      </g>
      <g transform="translate(385, 0)">
          <rect class="day" width="8" height="8" x="-24" y="0" fill="#ebedf0" data-count="0" data-date="2018-07-22"></rect>
          <rect class="day" width="8" height="8" x="-24" y="10" fill="#7bc96f" data-count="17" data-date="2018-07-23"></rect>
          <rect class="day" width="8" height="8" x="-24" y="20" fill="#c6e48b" data-count="4" data-date="2018-07-24"></rect>
          <rect class="day" width="8" height="8" x="-24" y="30" fill="#c6e48b" data-count="6" data-date="2018-07-25"></rect>
          <rect class="day" width="8" height="8" x="-24" y="40" fill="#c6e48b" data-count="6" data-date="2018-07-26"></rect>
          <rect class="day" width="8" height="8" x="-24" y="50" fill="#c6e48b" data-count="5" data-date="2018-07-27"></rect>
          <rect class="day" width="8" height="8" x="-24" y="60" fill="#ebedf0" data-count="0" data-date="2018-07-28"></rect>
      </g>
      <g transform="translate(396, 0)">
          <rect class="day" width="8" height="8" x="-25" y="0" fill="#ebedf0" data-count="0" data-date="2018-07-29"></rect>
          <rect class="day" width="8" height="8" x="-25" y="10" fill="#7bc96f" data-count="13" data-date="2018-07-30"></rect>
          <rect class="day" width="8" height="8" x="-25" y="20" fill="#c6e48b" data-count="1" data-date="2018-07-31"></rect>
          <rect class="day" width="8" height="8" x="-25" y="30" fill="#7bc96f" data-count="11" data-date="2018-08-01"></rect>
          <rect class="day" width="8" height="8" x="-25" y="40" fill="#7bc96f" data-count="14" data-date="2018-08-02"></rect>
          <rect class="day" width="8" height="8" x="-25" y="50" fill="#c6e48b" data-count="8" data-date="2018-08-03"></rect>
          <rect class="day" width="8" height="8" x="-25" y="60" fill="#ebedf0" data-count="0" data-date="2018-08-04"></rect>
      </g>
      <g transform="translate(407, 0)">
          <rect class="day" width="8" height="8" x="-26" y="0" fill="#ebedf0" data-count="0" data-date="2018-08-05"></rect>
          <rect class="day" width="8" height="8" x="-26" y="10" fill="#ebedf0" data-count="0" data-date="2018-08-06"></rect>
          <rect class="day" width="8" height="8" x="-26" y="20" fill="#7bc96f" data-count="11" data-date="2018-08-07"></rect>
          <rect class="day" width="8" height="8" x="-26" y="30" fill="#c6e48b" data-count="1" data-date="2018-08-08"></rect>
          <rect class="day" width="8" height="8" x="-26" y="40" fill="#7bc96f" data-count="15" data-date="2018-08-09"></rect>
          <rect class="day" width="8" height="8" x="-26" y="50" fill="#c6e48b" data-count="6" data-date="2018-08-10"></rect>
          <rect class="day" width="8" height="8" x="-26" y="60" fill="#ebedf0" data-count="0" data-date="2018-08-11"></rect>
      </g>
      <g transform="translate(418, 0)">
          <rect class="day" width="8" height="8" x="-27" y="0" fill="#ebedf0" data-count="0" data-date="2018-08-12"></rect>
          <rect class="day" width="8" height="8" x="-27" y="10" fill="#7bc96f" data-count="13" data-date="2018-08-13"></rect>
          <rect class="day" width="8" height="8" x="-27" y="20" fill="#c6e48b" data-count="3" data-date="2018-08-14"></rect>
          <rect class="day" width="8" height="8" x="-27" y="30" fill="#c6e48b" data-count="1" data-date="2018-08-15"></rect>
          <rect class="day" width="8" height="8" x="-27" y="40" fill="#c6e48b" data-count="6" data-date="2018-08-16"></rect>
          <rect class="day" width="8" height="8" x="-27" y="50" fill="#c6e48b" data-count="2" data-date="2018-08-17"></rect>
          <rect class="day" width="8" height="8" x="-27" y="60" fill="#ebedf0" data-count="0" data-date="2018-08-18"></rect>
      </g>
      <g transform="translate(429, 0)">
          <rect class="day" width="8" height="8" x="-28" y="0" fill="#ebedf0" data-count="0" data-date="2018-08-19"></rect>
          <rect class="day" width="8" height="8" x="-28" y="10" fill="#ebedf0" data-count="0" data-date="2018-08-20"></rect>
          <rect class="day" width="8" height="8" x="-28" y="20" fill="#7bc96f" data-count="14" data-date="2018-08-21"></rect>
          <rect class="day" width="8" height="8" x="-28" y="30" fill="#c6e48b" data-count="5" data-date="2018-08-22"></rect>
          <rect class="day" width="8" height="8" x="-28" y="40" fill="#c6e48b" data-count="4" data-date="2018-08-23"></rect>
          <rect class="day" width="8" height="8" x="-28" y="50" fill="#ebedf0" data-count="0" data-date="2018-08-24"></rect>
          <rect class="day" width="8" height="8" x="-28" y="60" fill="#c6e48b" data-count="6" data-date="2018-08-25"></rect>
      </g>
      <g transform="translate(440, 0)">
          <rect class="day" width="8" height="8" x="-29" y="0" fill="#c6e48b" data-count="6" data-date="2018-08-26"></rect>
          <rect class="day" width="8" height="8" x="-29" y="10" fill="#ebedf0" data-count="0" data-date="2018-08-27"></rect>
          <rect class="day" width="8" height="8" x="-29" y="20" fill="#c6e48b" data-count="7" data-date="2018-08-28"></rect>
          <rect class="day" width="8" height="8" x="-29" y="30" fill="#7bc96f" data-count="12" data-date="2018-08-29"></rect>
          <rect class="day" width="8" height="8" x="-29" y="40" fill="#7bc96f" data-count="10" data-date="2018-08-30"></rect>
          <rect class="day" width="8" height="8" x="-29" y="50" fill="#239a3b" data-count="21" data-date="2018-08-31"></rect>
          <rect class="day" width="8" height="8" x="-29" y="60" fill="#ebedf0" data-count="0" data-date="2018-09-01"></rect>
      </g>
      <g transform="translate(451, 0)">
          <rect class="day" width="8" height="8" x="-30" y="0" fill="#ebedf0" data-count="0" data-date="2018-09-02"></rect>
          <rect class="day" width="8" height="8" x="-30" y="10" fill="#7bc96f" data-count="17" data-date="2018-09-03"></rect>
          <rect class="day" width="8" height="8" x="-30" y="20" fill="#c6e48b" data-count="8" data-date="2018-09-04"></rect>
          <rect class="day" width="8" height="8" x="-30" y="30" fill="#196127" data-count="39" data-date="2018-09-05"></rect>
          <rect class="day" width="8" height="8" x="-30" y="40" fill="#7bc96f" data-count="10" data-date="2018-09-06"></rect>
          <rect class="day" width="8" height="8" x="-30" y="50" fill="#ebedf0" data-count="0" data-date="2018-09-07"></rect>
          <rect class="day" width="8" height="8" x="-30" y="60" fill="#ebedf0" data-count="0" data-date="2018-09-08"></rect>
      </g>
      <g transform="translate(462, 0)">
          <rect class="day" width="8" height="8" x="-31" y="0" fill="#ebedf0" data-count="0" data-date="2018-09-09"></rect>
          <rect class="day" width="8" height="8" x="-31" y="10" fill="#ebedf0" data-count="0" data-date="2018-09-10"></rect>
          <rect class="day" width="8" height="8" x="-31" y="20" fill="#196127" data-count="34" data-date="2018-09-11"></rect>
          <rect class="day" width="8" height="8" x="-31" y="30" fill="#196127" data-count="42" data-date="2018-09-12"></rect>
          <rect class="day" width="8" height="8" x="-31" y="40" fill="#7bc96f" data-count="11" data-date="2018-09-13"></rect>
          <rect class="day" width="8" height="8" x="-31" y="50" fill="#c6e48b" data-count="8" data-date="2018-09-14"></rect>
          <rect class="day" width="8" height="8" x="-31" y="60" fill="#ebedf0" data-count="0" data-date="2018-09-15"></rect>
      </g>
      <g transform="translate(473, 0)">
          <rect class="day" width="8" height="8" x="-32" y="0" fill="#ebedf0" data-count="0" data-date="2018-09-16"></rect>
          <rect class="day" width="8" height="8" x="-32" y="10" fill="#ebedf0" data-count="0" data-date="2018-09-17"></rect>
          <rect class="day" width="8" height="8" x="-32" y="20" fill="#196127" data-count="37" data-date="2018-09-18"></rect>
          <rect class="day" width="8" height="8" x="-32" y="30" fill="#ebedf0" data-count="0" data-date="2018-09-19"></rect>
          <rect class="day" width="8" height="8" x="-32" y="40" fill="#c6e48b" data-count="3" data-date="2018-09-20"></rect>
          <rect class="day" width="8" height="8" x="-32" y="50" fill="#7bc96f" data-count="16" data-date="2018-09-21"></rect>
          <rect class="day" width="8" height="8" x="-32" y="60" fill="#ebedf0" data-count="0" data-date="2018-09-22"></rect>
      </g>
      <g transform="translate(484, 0)">
          <rect class="day" width="8" height="8" x="-33" y="0" fill="#ebedf0" data-count="0" data-date="2018-09-23"></rect>
          <rect class="day" width="8" height="8" x="-33" y="10" fill="#ebedf0" data-count="0" data-date="2018-09-24"></rect>
          <rect class="day" width="8" height="8" x="-33" y="20" fill="#ebedf0" data-count="0" data-date="2018-09-25"></rect>
          <rect class="day" width="8" height="8" x="-33" y="30" fill="#c6e48b" data-count="7" data-date="2018-09-26"></rect>
          <rect class="day" width="8" height="8" x="-33" y="40" fill="#c6e48b" data-count="8" data-date="2018-09-27"></rect>
          <rect class="day" width="8" height="8" x="-33" y="50" fill="#c6e48b" data-count="8" data-date="2018-09-28"></rect>
          <rect class="day" width="8" height="8" x="-33" y="60" fill="#ebedf0" data-count="0" data-date="2018-09-29"></rect>
      </g>
      <g transform="translate(495, 0)">
          <rect class="day" width="8" height="8" x="-34" y="0" fill="#ebedf0" data-count="0" data-date="2018-09-30"></rect>
          <rect class="day" width="8" height="8" x="-34" y="10" fill="#ebedf0" data-count="0" data-date="2018-10-01"></rect>
          <rect class="day" width="8" height="8" x="-34" y="20" fill="#7bc96f" data-count="14" data-date="2018-10-02"></rect>
          <rect class="day" width="8" height="8" x="-34" y="30" fill="#7bc96f" data-count="16" data-date="2018-10-03"></rect>
          <rect class="day" width="8" height="8" x="-34" y="40" fill="#ebedf0" data-count="0" data-date="2018-10-04"></rect>
          <rect class="day" width="8" height="8" x="-34" y="50" fill="#239a3b" data-count="26" data-date="2018-10-05"></rect>
          <rect class="day" width="8" height="8" x="-34" y="60" fill="#ebedf0" data-count="0" data-date="2018-10-06"></rect>
      </g>
      <g transform="translate(506, 0)">
          <rect class="day" width="8" height="8" x="-35" y="0" fill="#ebedf0" data-count="0" data-date="2018-10-07"></rect>
          <rect class="day" width="8" height="8" x="-35" y="10" fill="#ebedf0" data-count="0" data-date="2018-10-08"></rect>
          <rect class="day" width="8" height="8" x="-35" y="20" fill="#c6e48b" data-count="6" data-date="2018-10-09"></rect>
          <rect class="day" width="8" height="8" x="-35" y="30" fill="#c6e48b" data-count="9" data-date="2018-10-10"></rect>
          <rect class="day" width="8" height="8" x="-35" y="40" fill="#c6e48b" data-count="1" data-date="2018-10-11"></rect>
          <rect class="day" width="8" height="8" x="-35" y="50" fill="#ebedf0" data-count="0" data-date="2018-10-12"></rect>
          <rect class="day" width="8" height="8" x="-35" y="60" fill="#ebedf0" data-count="0" data-date="2018-10-13"></rect>
      </g>
      <g transform="translate(517, 0)">
          <rect class="day" width="8" height="8" x="-36" y="0" fill="#ebedf0" data-count="0" data-date="2018-10-14"></rect>
          <rect class="day" width="8" height="8" x="-36" y="10" fill="#7bc96f" data-count="10" data-date="2018-10-15"></rect>
          <rect class="day" width="8" height="8" x="-36" y="20" fill="#c6e48b" data-count="2" data-date="2018-10-16"></rect>
          <rect class="day" width="8" height="8" x="-36" y="30" fill="#7bc96f" data-count="11" data-date="2018-10-17"></rect>
          <rect class="day" width="8" height="8" x="-36" y="40" fill="#c6e48b" data-count="9" data-date="2018-10-18"></rect>
          <rect class="day" width="8" height="8" x="-36" y="50" fill="#7bc96f" data-count="12" data-date="2018-10-19"></rect>
          <rect class="day" width="8" height="8" x="-36" y="60" fill="#ebedf0" data-count="0" data-date="2018-10-20"></rect>
      </g>
      <g transform="translate(528, 0)">
          <rect class="day" width="8" height="8" x="-37" y="0" fill="#ebedf0" data-count="0" data-date="2018-10-21"></rect>
          <rect class="day" width="8" height="8" x="-37" y="10" fill="#ebedf0" data-count="0" data-date="2018-10-22"></rect>
          <rect class="day" width="8" height="8" x="-37" y="20" fill="#c6e48b" data-count="2" data-date="2018-10-23"></rect>
          <rect class="day" width="8" height="8" x="-37" y="30" fill="#c6e48b" data-count="4" data-date="2018-10-24"></rect>
          <rect class="day" width="8" height="8" x="-37" y="40" fill="#ebedf0" data-count="0" data-date="2018-10-25"></rect>
          <rect class="day" width="8" height="8" x="-37" y="50" fill="#c6e48b" data-count="1" data-date="2018-10-26"></rect>
          <rect class="day" width="8" height="8" x="-37" y="60" fill="#ebedf0" data-count="0" data-date="2018-10-27"></rect>
      </g>
      <g transform="translate(539, 0)">
          <rect class="day" width="8" height="8" x="-38" y="0" fill="#ebedf0" data-count="0" data-date="2018-10-28"></rect>
          <rect class="day" width="8" height="8" x="-38" y="10" fill="#c6e48b" data-count="1" data-date="2018-10-29"></rect>
          <rect class="day" width="8" height="8" x="-38" y="20" fill="#c6e48b" data-count="7" data-date="2018-10-30"></rect>
          <rect class="day" width="8" height="8" x="-38" y="30" fill="#c6e48b" data-count="3" data-date="2018-10-31"></rect>
          <rect class="day" width="8" height="8" x="-38" y="40" fill="#c6e48b" data-count="7" data-date="2018-11-01"></rect>
          <rect class="day" width="8" height="8" x="-38" y="50" fill="#c6e48b" data-count="6" data-date="2018-11-02"></rect>
          <rect class="day" width="8" height="8" x="-38" y="60" fill="#ebedf0" data-count="0" data-date="2018-11-03"></rect>
      </g>
      <g transform="translate(550, 0)">
          <rect class="day" width="8" height="8" x="-39" y="0" fill="#ebedf0" data-count="0" data-date="2018-11-04"></rect>
          <rect class="day" width="8" height="8" x="-39" y="10" fill="#c6e48b" data-count="9" data-date="2018-11-05"></rect>
          <rect class="day" width="8" height="8" x="-39" y="20" fill="#c6e48b" data-count="9" data-date="2018-11-06"></rect>
          <rect class="day" width="8" height="8" x="-39" y="30" fill="#239a3b" data-count="26" data-date="2018-11-07"></rect>
          <rect class="day" width="8" height="8" x="-39" y="40" fill="#7bc96f" data-count="11" data-date="2018-11-08"></rect>
          <rect class="day" width="8" height="8" x="-39" y="50" fill="#ebedf0" data-count="0" data-date="2018-11-09"></rect>
          <rect class="day" width="8" height="8" x="-39" y="60" fill="#ebedf0" data-count="0" data-date="2018-11-10"></rect>
      </g>
      <g transform="translate(561, 0)">
          <rect class="day" width="8" height="8" x="-40" y="0" fill="#ebedf0" data-count="0" data-date="2018-11-11"></rect>
          <rect class="day" width="8" height="8" x="-40" y="10" fill="#c6e48b" data-count="1" data-date="2018-11-12"></rect>
          <rect class="day" width="8" height="8" x="-40" y="20" fill="#c6e48b" data-count="2" data-date="2018-11-13"></rect>
          <rect class="day" width="8" height="8" x="-40" y="30" fill="#c6e48b" data-count="2" data-date="2018-11-14"></rect>
          <rect class="day" width="8" height="8" x="-40" y="40" fill="#ebedf0" data-count="0" data-date="2018-11-15"></rect>
          <rect class="day" width="8" height="8" x="-40" y="50" fill="#ebedf0" data-count="0" data-date="2018-11-16"></rect>
          <rect class="day" width="8" height="8" x="-40" y="60" fill="#ebedf0" data-count="0" data-date="2018-11-17"></rect>
      </g>
      <g transform="translate(572, 0)">
          <rect class="day" width="8" height="8" x="-41" y="0" fill="#c6e48b" data-count="2" data-date="2018-11-18"></rect>
          <rect class="day" width="8" height="8" x="-41" y="10" fill="#c6e48b" data-count="6" data-date="2018-11-19"></rect>
          <rect class="day" width="8" height="8" x="-41" y="20" fill="#c6e48b" data-count="9" data-date="2018-11-20"></rect>
      </g>
      <text x="11" y="-8" class="month">Nov</text>
      <text x="31" y="-8" class="month">Dec</text>
      <text x="81" y="-8" class="month">Jan</text>
      <text x="121" y="-8" class="month">Feb</text>
      <text x="161" y="-8" class="month">Mar</text>
      <text x="201" y="-8" class="month">Apr</text>
      <text x="251" y="-8" class="month">May</text>
      <text x="291" y="-8" class="month">Jun</text>
      <text x="331" y="-8" class="month">Jul</text>
      <text x="381" y="-8" class="month">Aug</text>
      <text x="421" y="-8" class="month">Sep</text>
      <text x="471" y="-8" class="month">Oct</text>
      <text x="511" y="-8" class="month">Nov</text>
    <text text-anchor="start" class="wday" dx="-12" dy="8" style="display: none;">Sun</text>
    <text text-anchor="start" class="wday" dx="-12" dy="17">Mon</text>
    <text text-anchor="start" class="wday" dx="-12" dy="32" style="display: none;">Tue</text>
    <text text-anchor="start" class="wday" dx="-12" dy="37">Wed</text>
    <text text-anchor="start" class="wday" dx="-12" dy="57" style="display: none;">Thu</text>
    <text text-anchor="start" class="wday" dx="-12" dy="57">Fri</text>
    <text text-anchor="start" class="wday" dx="-12" dy="81" style="display: none;">Sat</text>
  </g>
</svg>`

		Context("when real endpoint", func() {
			Context("when valid username", func() {
				It("expects no error", func() {
					// confirmation real endpoint
					contributions, err := client.FetchContributions("hiko1129")
					Expect(contributions).NotTo(BeEmpty())
					Expect(err).NotTo(HaveOccurred())
				})
			})

			// Context("when invalid username", func() {
			// 	It("expects error", func() {
			// 		httpmock.Activate()
			// 		defer httpmock.DeactivateAndReset()
			// 		httpmock.RegisterResponder("GET", "https://github.com/users/hiko1129/contributions", httpmock.NewStringResponder(200, html))

			// 		contributions, err := client.FetchContributions("-01-01-01-")
			// 		Expect(contributions).To(BeEmpty())
			// 		Expect(err).To(HaveOccurred())
			// 	})
			// })
		})

		Context("when fake endpoint", func() {
			BeforeEach(func() {
				httpmock.Activate()
			})

			AfterEach(func() {
				httpmock.DeactivateAndReset()
			})

			Context("when valid username", func() {
				It("expects no error", func() {
					httpmock.RegisterResponder("GET", "https://github.com/users/hiko1129/contributions", httpmock.NewStringResponder(200, html))
					contributions, err := client.FetchContributions("hiko1129")
					Expect(contributions).NotTo(BeEmpty())
					Expect(err).NotTo(HaveOccurred())

					contributions, err = client.FetchContributions("hiko1129")
					Expect(contributions).NotTo(BeEmpty())
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when invalid username", func() {
				It("expects 404 error", func() {
					httpmock.RegisterResponder("GET", "https://github.com/users/abcdefghijklmn/contributions", httpmock.NewStringResponder(404, "not found"))
					contributions, err := client.FetchContributions("abcdefghijklmn")
					Expect(contributions).To(BeEmpty())
					Expect(err.Error()).To(ContainSubstring("status code error: 404"))
				})

				It("expects error", func() {
					contributions, err := client.FetchContributions("abcdefghijklmn")
					Expect(contributions).To(BeEmpty())
					Expect(err.Error()).To(ContainSubstring("no responder found"))
				})
			})
		})
	})
})
