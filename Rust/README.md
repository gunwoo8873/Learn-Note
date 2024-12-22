# Rust-Learn
> 각 하단 아래에 Comment 형식으로 요약문 작성
* Manual
    ```bash
    cargo build
    cargo run --bin [bin_name]
    ```
    ```bash
    cargo test
    ```
# Bin Name list
| Name                                         | Update Date |
| -------------------------------------------- | ----------- |
| 1. [Variable](./src/bin/Variable.rs)         | 24. 12. 21  |
| 2. [Control Flow](./src/bin/Control_Flow.rs) | 24. 12. 21  |
| 3. [Ownership](./src/bin/Ownership.rs)       | 24. 12. 21  |
| 4. [Thread](./src/bin/Thread.rs)             | 24. 12. 21  |
| 5. [Options](./src/bin/Options.rs)           | 24. 12. 21  |
| 6. [Collection](./src/bin/Options.rs)        | 24. 12. 21  |
| 7. [Trait](./src/bin/Trait.rs)               | 24. 12. 22  |
| 8. [Atomic](./src/bin/Atomic.rs)             | 24. 12. 21  |
| 9. [Panic](./src/bin/Panic.rs)               | 24. 12. 21  |
| 10. [Pattern](./src/bin/Pattern.rs)          | 24. 12. 21  |
> Noti : Grap file CMD : cargo run --bin grap -- test sample.txt


# Build Check
```bash
cargo clean
```

```bash
cargo build
```

```bash
time cargo check
```

# Toolchain Change
> Rust Toolchain for Stable, Beta, Nightly to need Installed

```bash
cargo +[version] test
```

```bash
cargo +stable test
cargo +beta test
cargo +nightly test
```

# Crate Lib Version Check
| Operator | EX      | Min      | Max     | Update Check |
| -------- | ------- | -------- | ------- | ------------ |
| ^        | ^ 1.1.0 | >= 1.1.0 | < 2.0.0 | Yes          |
| ~        | ~ 1.2.3 | >= 1.2.3 | < 1.3.0 | Yes          |