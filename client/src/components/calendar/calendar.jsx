import { createSignal, onMount } from 'solid-js';
import FullCalendar from '@fullcalendar/react';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import '../../../../app/internal/public/css/pages/calendar/calendar.css';

function MyCalendar() {
  const [view, setView] = createSignal('dayGridMonth');

  onMount(() => {
    const script = document.createElement('script');
    script.src = '/calendar/extensions/index.global.min.js';
    script.onload = () => {
      toggleVisibility();
      window.addEventListener('resize', toggleVisibility);
    };
    document.head.appendChild(script);
  });

  function showCalendar(gridCall) {
    const calendarEl = document.getElementById('calendar');
    if (calendarEl) {
      const calendar = new FullCalendar.Calendar(calendarEl, {
        plugins: [dayGridPlugin, timeGridPlugin, multiMonthPlugin],
        dateClick: function (info) {
          showCalendar('dayGridMonth');
        },
        headerToolbar: {
          left: 'prev,next today',
          center: 'title',
          right: 'timeGridDay,timeGridWeek,dayGridMonth,multiMonthYear',
        },
        initialView: gridCall,
        fixedWeekCount: false,
      });
      calendar.render();
    } else {
      console.error('Calendar element not found');
    }
  }

  function toggleVisibility() {
    if (window.innerWidth <= 768) {
      setView('timeGridWeek');
      showCalendar('timeGridWeek');
    } else {
      setView('dayGridMonth');
      showCalendar('dayGridMonth');
    }
  }

  return (
    <div>
      <h1>Calendar</h1>
      <center>
        <div id="calendar"></div>
      </center>
    </div>
  );
}

export default MyCalendar;
