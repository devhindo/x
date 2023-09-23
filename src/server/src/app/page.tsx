import styles from './page.module.css'

export default function Home() {
  return (
    <div>
      <main className={styles.main}>
        X CLI
        <p>tweet from terminal and stuff</p>
        <fieldset>
          <legend>Installation</legend>
          <ul>
            <li></li>
          </ul>
        </fieldset>
        <fieldset>
          <legend>Authorization</legend>
          <ul>
            <li>Run <span className={styles.cmd}>x auth</span></li>
            <li>Then <span className={styles.cmd}>x auth -v</span> or <span className={styles.cmd}>x auth --validate</span></li>
          </ul>
        </fieldset>
        <fieldset>
          <legend>commands</legend>
          <ul>
            <li>tweet: <span className={styles.cmd}>x -t "hello, xcli!"</span> or <span className={styles.cmd}>x --tweet "hello, xcli!"</span></li>
          </ul>
        </fieldset>
      </main>
      <h1><a href="https://github.com/devhindo/x" target='_blank'>source code</a></h1>
    </div>
  )
}
