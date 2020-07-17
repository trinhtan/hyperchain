## 1. Orgs config

Lưu trong file config/orgsconfig.yaml

Gồm có:
- domain: tên domain
- orgs: Các orgs:
 + Mỗi Org có tên, số peer, cpu, ram, storage (đơn vị Gi)

File generate_networkconfig.py sẽ đọc dữ liệu của file này và generate ra thư mục:

projects/${domain}/network/deploy-kubernetes/kubernetes/

và các file scripts.sh, network.sh, chaincode.sh

- File scripts.sh để up các pvc, pv, pods, deployments, services cho các thành phần của fabric network, excute các file network.sh, chaincode.sh trong pod fabric-${domain}-tools

- File network.sh để tạo các cấu hình cần thiết cho network như thư mục crypto-config, channel-artifacts

- File chaincode.sh để tạo channel, join các peer vào channel, install chaincode và instantiate chaincode

### 2. Chaincode config

- Lưu trong file config/orgsconfig.yaml

- domain: domain này phải giống với domain của orgs config

- objects: các object

- Mỗi object có properties, tên của tất cả property phải viết hoa chữ cái đầu, type là string, không cần định nghĩa trường Id cho object vì sẽ lấy mặc định là {{obj['name']|capitalize}}Id. Ví dụ object course có trường Id là CourseId

- Mỗi object sẽ có các hàm măc định là GetAll và Get theo Id, chẳng hạn như GetAllCourse(), GetCourse(CourseId)

#### 2.1 Create:
- function_name có thể do người dùng đặt hoặc mặc định là Create{{obj['name']|capitalize}}. Ví dụ CreaetCourse. Tham số truyền vào là tất cả properties đã được định nghĩa.

- msp: 
    + Có thể là null, trong chaincode sẽ không kiểm tra identity của user gọi invoke chaincode 
    + Hoặc là mảng các tên org, các org này phải là thuộc các org trong file orgsconfig.yaml. Ví dụ
    ```yaml
        msp:
            - academy
    ```
    nghĩa là chỉ có org academy mới có quyển thực hiện CreateCourse

#### 2.2 Update:
- function_name có thể do người dùng đặt hoặc mặc định là Update{{obj['name']|capitalize}}. Ví dụ UpdateCourse. Tham số truyền vào là CourseId và tất cả properties.

- msp:
    + Có thể là null, trong chaincode sẽ không kiểm tra identity của user gọi invoke chaincode 
    + Hoặc là mảng các tên org, các org này phải là thuộc các org trong file orgsconfig.yaml. Ví dụ
    ```yaml
        msp:
            - academy
    ```
    nghĩa là chỉ có org academy mới có quyển thực hiện UpdateCourse

#### 2.4 Delete - Tạm thời pending

#### 2.5 Inclusion - Is Inclused

Ví dụ Course inclusion subject

```yaml
inclusion:
    - child_obj_name: subject
    child_obj_name_many: subjects
    field: Subjects
        add:
            function_name: AddSubjectToCourse
            msp: Null
            # msp: 
            # - academy
            # - student
        remove:
            function_name: RemoveSubjectFromCourse
            msp: Null
            # msp:
            # - academy
            # - student
```
- child_obj_name: Tên của obj mà nó inclusion
- child_obj_name_many: Tên số nhiều của obj mà nó nó inclusion
- field: Tên trường mà nó dùng để lưu sự inclusion. Khi đó trong nó sẽ có thêm một trường là []string để lưu.

Ví dụ trong Course sẽ có trường Subjects []string để lưu các Subjects mà nó có. Một Course có thể có nhiều Subject và một subject có thể thuộc nhiều Course. Chẳng hạn như một ngành học có nhiều môn, và các môn ấy có thể thuộc nhiều ngành học.

phương thức add và remove sẽ có:
 - function_name: tên function trong chaincode, có thể  là mặc định theo khuôn AddAToB, RemoveAFromB. Tham số truyền vào lần lượt là : AddSubjectToCourse(CourseId, SubjectId)m RemoveSubjectFromCourse(CourseId, SubjectId)
 - msp: tên các tổ chức có quyền thực hiện function đó 

và trong Subject sẽ phải có một định nghĩa is_inclused để đối xứng lại với inclusion của course.

```yaml
is_inclused:
    - parent_obj_name: course
    parent_obj_name_many: courses
    parent_field: Subjects
```

- Có thể tạo một CourseA, tạo một SubjectB, sau đó mới thực hiện AddSubjectToCourse(CousreA, SubjectB)

- Với sự inclusion, sẽ có thêm các hàm mặc định, chẳng hạn như GetSubjectsOfCourse(CourseId)
    
#### 2.6 Dependence - Is dependenced

Ví dụ Subject dependence class

```yaml
dependence: 
    - child_obj_name: class
    child_obj_name_many: classes
    field: Classes
```
Khi đó trong Subject sẽ có thêm một trường là: Classes []string

Và có thêm một hàm mặc định GetClassesOfSubject(SubjectId)

Class phải có một định nghĩa is_dependenced đê đối xứng lại

```yaml
is_dependenced:
    - parent_obj_name: subject
    parent_obj_name_many: subjects
    parent_field: Classes
    field: SubjectId
```
trong Class sẽ có thêm một trường là SubjectId để lưu subject mà nó thuộc về.

Và hàm CreateClass sẽ có thêm tham số SubjectId. Nghĩa là khi tạo class bắt buộc phải chỉ định luôn Subject mà nó thuộc về

Không thể tạo class mà không biết nó thuộc subject nào

#### 2.7 Owner - Is owned

Ví dụ teacher owner class:

```yaml
owner:
    - child_obj_name: class
    child_obj_name_many: classes
    child_obj_field: TeacherId
    field: Classes
    add: 
        function_name: AssignClassToTeacher
        msp: Null
    remove:
        function_name: UnassignClassFromTeacher
        msp: Null
        # msp:
        # - AcademyMSP
        # - StudentMSP
    change:
        function_name: ChangeTeacher
        msp: Null
        # msp:
        # - AcademyMSP
        # - StudentMSP
```
Teacher sở hữu một class, Khi đó trong teacher sẽ có một trường: Classes []string để lưu các class mà teahcher đó sở hữu. Trong class sẽ có một trường là TeacherId để lưu techer mà nó thuộc về. Nghĩa là ta sẽ tạo một ClassA, một TeacherB sau đó mới thực hiện gán quyền sở hửu class cho teacher bằng hàm AssignClassToTeacher(TeacherId, ClassId)

- add: AssignClassToTeacher(TeacherId, ClassId)
- remove: UnassignClassFromTeacher(TeacherId, ClassId)
- change: ChangeTeacher(ClassId, newTeahcerId)

- Sẽ có thêm một hàm mặc định là GetClassesOfTeacher(TeacherId) để get tất cả classes mà teacher đấy sở hữu.

Phía class phải có một định nghĩa is_owned để đối xứng lại.

```yaml
 is_owned:
      - parent_obj_name: teacher
        parent_obj_name_many: teachers
        parent_field: Classes
        field: TeacherId
```
- Sẽ có thêm một hàm mặc định là GetTeacherOfClass(ClassId) để get teacher của class đấy


### 2.8 Match - Is matched

Ví dụ student match với course. Một student có thể đăng ký nhiều course, một course sẽ có nhiều students

```yaml
match:
    - des_obj_name: course
    des_obj_name_many: courses
    des_field: Students
    field: Courses
    add:
        function_name: StudentResgisterCourse
        msp: Null
        # msp:
        # - student
    remove:
        function_name: StudentCancelCourse
        msp: Null
        # msp:
        # - student
```
Khi đó trong student sẽ có thêm trường Courses []string và trong Course sẽ có thêm trường Students []string

- add: StudentRegisterCourse(StudentId, CourseId)
- remove: StudentCancleCourse(StudentId, CourseId)

Khi đó sẽ có thêm một hàm mặc định GetCoursesOfStudent(StudentId)

Bên course phải có một định nghĩa is_matched để đối xứng lại

```yaml
is_matched: 
    - sou_obj_name: student
    sou_obj_name_many: students
    sou_field: Courses
    field: Students
```
Có thêm một hàm mặc định GetStudentsOfCourse(CourseId)

### 3. Cách test

- Cài minikube 

- Start minkube
```bash
minikube start
```

#### 3.1 Work flow

Bước 1: Up network

```bash
./up_network ${PROJECT_NAME}
```

Ví dụ
```bash
./up_networl certificate
```

Bước 2: Install chaincode lên network đã được up

```bash
./install_chaincode certificate
```

Bước 3: Sau khi chaincode được install nếu muốn cập nhật chaincode mới, thì upgrade chaincode
- Tăng version thành 2.0, 3.0 trong chaincodeconfig.yaml
- Gọi lệnh
```bash
./upgrade_chaincode.sh certificate
```

Bước 4: Down network
```bash
./down_network.sh certificate
```

### 4. Tổng kết

Với config như chaincodeconfig.yaml sẽ sinh ra file certificate.go