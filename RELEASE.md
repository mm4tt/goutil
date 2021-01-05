# Release

## ChangeLog

#### 0.2.0
- Enhanced the PriorityQueue by replacing the priority function with the more
  generic less function.

#### 0.1.1
- Fixed a bug in the BitSet.
  
#### 0.1.0
- Initial version

## How to release

1. Create new tag, e.g. new minor release
    ```
    git tag v0.3.0
    ```

2. Push the tag
    ```
    git push origin v0.3.0
    ```

3. Updating clients
    ```
    go get github.com/mm4tt/goutil@v0.3.0
    ```
