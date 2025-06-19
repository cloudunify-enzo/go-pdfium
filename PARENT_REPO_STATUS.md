# Status Update: Parent Repository (klippa-app/go-pdfium)

This document summarizes recent notable changes in the parent repository, `klippa-app/go-pdfium`. This fork, `cloudunify-enzo/go-pdfium`, should consider these upstream developments.

## Key Highlights:

*   **PDFium Library Updated to Version 7019:**
    *   The core PDFium C++ library, which this Go binding wraps, has been updated to version 7019. This is a significant upgrade and likely includes new features, performance improvements, and bug fixes from the PDFium project.
    *   New methods from this PDFium version have been implemented in `klippa-app/go-pdfium`.
    *   Tests have been updated and extended to cover these new methods and changes in PDFium behavior.

*   **Go Version Support Policy:**
    *   The parent repository has explicitly updated its Go version support policy to align with the official Go support policy (supporting the last two major Go versions).

*   **Concurrency Improvements:**
    *   Fixes and improvements have been made to instance locking mechanisms, particularly around instance creation. This should enhance stability and predictability in multi-threaded scenarios.

*   **WebAssembly (WASM) Updates:**
    *   The WebAssembly build of the library has received updates, ensuring it stays current with other improvements.

*   **Dependency Management:**
    *   Regular updates to various Go module dependencies, including `golang.org/x/net`, `golang.org/x/text`, `github.com/onsi/ginkgo`, `github.com/stretchr/testify`, `github.com/hashicorp/go-plugin`, and `github.com/tetratelabs/wazero`, have been merged.

## Details from Recent Commits (Late Feb - Early March 2025):

*   **PDFium 7019 Upgrade (Feb 19-20, 2025):**
    *   Commits `dabee1f`, `ab38b15`, `5eb232b`, `11ee051`, `9fd3233`, `d00d3e2`, `45fae4c`, and `3516e50` reflect the work to update to PDFium 7019, implement new methods, and adjust tests.
*   **Go Version Policy (Mar 6, 2025):**
    *   Commit `926e953` and merge `5259e4c` relate to updating the Go version support.
*   **Instance Locking (Feb 25, 2025):**
    *   Commits `34b12d9` and merge `aa6084c` (from NickHilton/lock-on-instance) address instance creation locking.
*   **WebAssembly Build (Feb 19, 2025):**
    *   Commit `afb782a` specifically mentions updating the WebAssembly build.

This summary should help in understanding the recent trajectory of the parent repository. It's recommended to review these changes in detail if planning to merge upstream updates into this fork.
