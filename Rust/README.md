# Rust-Learn
> 각 하단 아래에 Comment 형식으로 요약문 작성
* Manual
    ```bash
    cargo build
    cargo run --bin [bin_name]
    ```
# Bin Name list
| Name          | Update Date |
| ------------- | ----------- |
| 1. Variable   | 24. 11. 13  |
| 2. Datatype   | 24. 11. 31  |
| 3. Operator   | 24. 11. 15  |
| 4. Statement  | 24. 11. 15  |
| 5. Ownership  | 24. 11. 13  |
| 6. Thread     | 24. 11. 06  |
| 7. Method     | 24. 11. 14  |
| 8. Options    | 24. 11. 15  |
| 9. Collection | 24. 11. 16  |

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