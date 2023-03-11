require("expose-loader?exposes=$,jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");
import { setUpFullCalendar } from './calendar/calendar.js';
import MicroModal from 'micromodal';



document.addEventListener("DOMContentLoaded", function () {
  MicroModal.init();

  const onDateClick = info => {
    console.log(`This is the info`)
    console.log(info)
    MicroModal.show('add-event')
  }



  const calendarContainer = document.getElementById("calendar");
  const calendar = setUpFullCalendar(calendarContainer, {
    dateClick: onDateClick
  });


});
