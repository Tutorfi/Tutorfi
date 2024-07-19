import FullCalendar from '@fullcalendar/react'

function CalendarPage() {
const [view, setView] = createSignal("dayGridMonth");

onMount(() => {
    const script = document.createElement('script');
    script.src = '/extensions/index.global.min.js';
    script.onload = () => {
    toggleVisibility();
    window.addEventListener("resize", toggleVisibility);
    };
    document.head.appendChild(script);
});

function showCalendar(gridCall) {
    const calendarEl = document.getElementById("calendar");
    const calendar = new FullCalendar.Calendar(calendarEl, {
    dateClick: function (info) {
        showCalendar("dayGridMonth");
    },
    headerToolbar: {
        left: "prev,next today",
        center: "title",
        right: "timeGridDay,timeGridWeek,dayGridMonth,multiMonthYear",
    },
        initialView: gridCall,
        fixedWeekCount: false,
    });
    calendar.render();
}

function toggleVisibility() {
    if (window.innerWidth <= 768) {
        
        setView("timeGridWeek");
        showCalendar("timeGridWeek");
    } 
    else {
        setView("dayGridMonth");
        showCalendar("dayGridMonth");
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

export default CalendarPage;
