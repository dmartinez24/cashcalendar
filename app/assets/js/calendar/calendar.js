import { Calendar } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import listPlugin from '@fullcalendar/list';
import interactionPlugin from '@fullcalendar/interaction';
import enLocale from '@fullcalendar/core/locales/en-gb';



const defaultOptions = {
  plugins: [dayGridPlugin, timeGridPlugin, listPlugin, interactionPlugin],
  initialView: 'dayGridMonth',
  locale: enLocale,
  selectable: true,
  nowIndicator: true,
  weekends: true,
  editable: true,
  headerToolbar: {
    left: 'prev, next, today',
    center: 'title',
    right: 'dayGridMonth, timeGridWeek, listWeek',
  },
  select: (info) => {
    console.log('%c SELECT INFO', 'color: lightpink');
    console.log(info);
  },
  events: [
    {
      // this object will be "parsed" into an Event Object
      title: 'Arya Stark', // a property!
      start: new Date(), // a property!
      display: 'block', // list-item, background
      allDay: true,
    },
  ],
};

export function setUpFullCalendar(element, options = {}) {
  const opts = {
    ...defaultOptions,
    ...options
  }
  const calendar = new Calendar(element, opts);

  calendar.render();

  return calendar;
}
