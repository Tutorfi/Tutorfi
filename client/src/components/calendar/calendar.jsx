import { createSignal } from 'solid-js';
import styles from './calendar.module.css';

function Calendar() {
const [currentDate, setCurrentDate] = createSignal(new Date());
const today = new Date();

const daysOfWeek = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];

const getDaysInMonth = (year, month) => {
    return new Date(year, month + 1, 0).getDate();
};

const generateCalendar = () => {
    const date = currentDate();
    const year = date.getFullYear();
    const month = date.getMonth();

    const firstDayOfMonth = new Date(year, month, 1).getDay();
    const daysInMonth = getDaysInMonth(year, month);

    const calendar = [];
    let day = 1;
    let rows;

    if ((firstDayOfMonth === 5 && daysInMonth === 31) || (firstDayOfMonth === 6 && daysInMonth >= 30)) {
    rows = 6;
    } else {
    rows = 5;
    }

    for (let i = 0; i < rows; i++) {
    const week = [];
    for (let j = 0; j < 7; j++) {
        const dayClass = `day${daysOfWeek[j].slice(0, 3)}`;
        if (i === 0 && j < firstDayOfMonth) {
        week.push(<td class={`${styles.dayOther} ${styles.days}`}></td>);
        } else if (day > daysInMonth) {
        week.push(<td class={`${styles.dayOther} ${styles.days}`}></td>);
        } else {
        const cellDate = new Date(year, month, day);
        let dayStateClass = '';
        if (cellDate.toDateString() === today.toDateString()) {
            dayStateClass = styles.dayToday;
        } else if (cellDate < today) {
            dayStateClass = styles.dayPast;
        } else {
            dayStateClass = styles.dayFuture;
        }

        week.push(
            <td class={dayStateClass}>
                <div class={styles.days}>
                    {day}
                </div>
                <div class={styles.details}>

                </div>
            </td>
        );
        day++;
        }
    }
    calendar.push(<tr>{week}</tr>);
    }

    return calendar;
};

const goToPreviousMonth = () => {
    const date = new Date(currentDate().setMonth(currentDate().getMonth() - 1));
    setCurrentDate(date);
};

const goToNextMonth = () => {
    const date = new Date(currentDate().setMonth(currentDate().getMonth() + 1));
    setCurrentDate(date);
};

const dayView = () => {
    
}

const weekView = () => {
    
}

const monthView = () => {
    
}

const yearView = () => {
    
}

return (
    <div class={styles['calendar-container']}>
        <div class={styles.header}>
            <div class={styles.monthButtons}>
                <button name='previous month' onClick={goToPreviousMonth}>&#9664;</button> {/* Left arrow */}
                <button name='next month' onClick={goToNextMonth}>&#9654;</button> {/* Right arrow */}
            </div>
            <h2>{currentDate().toLocaleDateString('en-US', { month: 'long', year: 'numeric' })}</h2>
            <div class={styles.modes}>
                <button name='day view' onClick={dayView}>day</button>
                <button name='week view' onClick={weekView}>week</button>
                <button name='month view' onClick={monthView}>month</button>
                <button name='year view' onClick={yearView}>year</button>
            </div>
        </div>
        <table class={styles.calendar}>
            <thead>
            <tr>
                {daysOfWeek.map(day => (
                <th>{day}</th>
                ))}
            </tr>
            </thead>
            <tbody>
            {generateCalendar()}
            </tbody>
        </table>
    </div>
);
}

export default Calendar;
