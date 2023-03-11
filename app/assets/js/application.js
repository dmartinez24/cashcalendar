require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");
const { setUpFullCalendar } = require("./calendar/calendar");



document.addEventListener("DOMContentLoaded", function () {
  const calendarContainer = document.getElementById("calendar");

  setUpFullCalendar(calendarContainer);
});
