# Rust-Learn
> 각 하단 아래에 Comment 형식으로 요약문 작성
* Manual
    ```bash
    cargo build
    cargo run --bin [bin_name]
    ```
# Bin Name list
| Name           | Update Date |
|----------------|-------------|
| 1. variable    | 24. 11. 13  |
| 2. datatype    | 24. 11. 31  |
| 3. operator    | 24. 11. 15  |
| 4. function    | 24. 11. 15  |
| 5. branches    | 24. 11. 15  |
| 6. ownership   | 24. 11. 13  |
| 7. thread      | 24. 11. 06  |
| 8. method      | 24. 11. 14  |
| 9. options     | 24. 11. 15  |
| 10. module     | 24. 11. 15  |
| 11. collection | 24. 11. 16  |

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
|----------|---------|----------|---------|--------------|
| ^        | ^ 1.1.0 | >= 1.1.0 | < 2.0.0 | Yes          | 
| ~        | ~ 1.2.3 | >= 1.2.3 | < 1.3.0 | Yes          |