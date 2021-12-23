# Release

## ChangeLog

#### 0.6.0
- Migrated to go1.18 and migrated the priority-queue to generics.

#### 0.5.0
- Introduced the async library.

#### 0.4.2
- Fixed a bug in assert.NotNil.

#### 0.4.1
- Fixed a bug in assert package.

#### 0.4.0
- Added assert.NotNil method.

#### 0.3.1
- Got rid of the external goerrors dep.

#### 0.3.0
- Introduced the "assert" library.

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

