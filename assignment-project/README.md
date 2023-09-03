# Assignment Project
This is assignment project for Golang for Women class. Details on the project is explained in Google Classroom (not public).

## API
| METHOD | URL                   | Description                   |
|--------|-----------------------|-------------------------------|
| GET    | `/students`           | Get all students              |
| POST   | `/students`           | Create new student            |
| GET    | `/students/studentID` | Get student of specific ID    |
| PUT    | `/students/studentID` | Update student of specific ID |
| DELETE | `/students/studentID` | Delete student of specific ID |

## Notes
Constraint `ON DELETE CASCADE` is set to `scores` table to avoid orphans (which means when a student is deleted, its corresponding record in `scores` table is also deleted). Since gorm `AutoMigrate` doesn't automatically updating constraint when you already have an existing table, you need to drop the table first before running the `AutoMigrate` to make sure **API delete student** works correctly.
