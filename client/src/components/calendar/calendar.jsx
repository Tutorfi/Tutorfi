import { createSignal } from 'solid-js';
import styles from './calendar.module.css';

function Calendar() {
  const [currentDate, setCurrentDate] = createSignal(new Date());

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

    for (let i = 0; i < 6; i++) {
      const week = [];
      for (let j = 0; j < 7; j++) {
        if (i === 0 && j < firstDayOfMonth) {
          week.push(<td class={styles.empty}></td>);
        } else if (day > daysInMonth) {
          week.push(<td class={styles.empty}></td>);
        } else {
          week.push(<td>{day}</td>);
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

return (
    <>
        <h1>Calendar</h1>
        <div class={styles.header}>
            <button onClick={goToPreviousMonth}>Previous</button>
            <h2>{currentDate().toLocaleDateString('en-US', { month: 'long', year: 'numeric' })}</h2>
            <button onClick={goToNextMonth}>Next</button>
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
    </>
);
}

export default Calendar;
