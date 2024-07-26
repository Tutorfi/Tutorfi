import { createSignal, onMount } from 'solid-js';
import styles from './calendar.module.css';

function Calendar() {
  const [days, setDays] = createSignal([]);

  onMount(() => {
    const date = new Date();
    const year = date.getFullYear();
    const month = date.getMonth();

    const firstDay = new Date(year, month, 1).getDay();
    const lastDate = new Date(year, month + 1, 0).getDate();

    const dayArray = [];
    for (let i = 0; i < firstDay; i++) {
      dayArray.push("");
    }
    for (let i = 1; i <= lastDate; i++) {
      dayArray.push(i);
    }
    setDays(dayArray);
  });

  return (
    <>
      <h1>Calendar</h1>
      <table class={styles.calendar}>
        <thead>
          <tr>
            <th>Sunday</th>
            <th>Monday</th>
            <th>Tuesday</th>
            <th>Wednesday</th>
            <th>Thursday</th>
            <th>Friday</th>
            <th>Saturday</th>
          </tr>
        </thead>
        <tbody>
          {Array.from({ length: Math.ceil(days().length / 7) }, (_, weekIndex) => (
            <tr>
              {Array.from({ length: 7 }, (_, dayIndex) => {
                const day = days()[weekIndex * 7 + dayIndex];
                return <td>{day}</td>;
              })}
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
}

export default Calendar;
