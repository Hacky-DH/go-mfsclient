package moosefs

const (
    VERSION_ANY       = 0
    CRC_POLY          = 0xEDB88320
    MFS_ROOT_ID       = 1
    MFS_NAME_MAX      = 255
    MFS_MAX_FILE_SIZE = 0x20000000000
)

// 1.6.20
const VERSION = uint32(0x01F1F5)

const GETDIR_FLAG_WITHATTR = 0x01

//type for readdir command 
const (
    TYPE_FILE      = 'f'
    TYPE_SYMLINK   = 'l'
    TYPE_DIRECTORY = 'd'
    TYPE_FIFO      = 'q'
    TYPE_BLOCKDEV  = 'b'
    TYPE_CHARDEV   = 'c'
    TYPE_SOCKET    = 's'
    TYPE_TRASH     = 't'
    TYPE_RESERVED  = 'r'
    TYPE_UNKNOWN   = '?'
)

// status code
const (
    STATUS_OK = iota // OK

    ERROR_EPERM       // Operation not permitted
    ERROR_ENOTDIR     // Not a directory
    ERROR_ENOENT      // No such file or directory
    ERROR_EACCES      // Permission denied
    ERROR_EEXIST      // File exists
    ERROR_EINVAL      // Invalid argument
    ERROR_ENOTEMPTY   // Directory not empty
    ERROR_CHUNKLOST   // Chunk lost
    ERROR_OUTOFMEMORY // Out of memory

    ERROR_INDEXTOOBIG    // Index too big
    ERROR_LOCKED         // Chunk locked
    ERROR_NOCHUNKSERVERS // No chunk servers
    ERROR_NOCHUNK        // No such chunk
    ERROR_CHUNKBUSY      // Chunk is busy
    ERROR_REGISTER       // Incorrect register BLOB
    ERROR_NOTDONE        // None of chunk servers performed requested operation
    ERROR_NOTOPENED      // File not opened
    ERROR_NOTSTARTED     // Write not started

    ERROR_WRONGVERSION   // Wrong chunk version
    ERROR_CHUNKEXIST     // Chunk already exists
    ERROR_NOSPACE        // No space left
    ERROR_IO             // IO error
    ERROR_BNUMTOOBIG     // Incorrect block number
    ERROR_WRONGSIZE      // Incorrect size
    ERROR_WRONGOFFSET    // Incorrect offset
    ERROR_CANTCONNECT    // Can't connect
    ERROR_WRONGCHUNKID   // Incorrect chunk id
    ERROR_DISCONNECTED   // Disconnected
    ERROR_CRC            // CRC error
    ERROR_DELAYED        // Operation delayed
    ERROR_CANTCREATEPATH // Can't create path

    ERROR_MISMATCH // Data mismatch

    ERROR_EROFS        // Read-only file system
    ERROR_QUOTA        // Quota exceeded
    ERROR_BADSESSIONID // Bad session id
    ERROR_NOPASSWORD   // Password is needed
    ERROR_BADPASSWORD  // Incorrect password

    ERROR_MAX
)

// flags: "flags" fileld in "CUTOMA_FUSE_AQUIRE"
const (
    WANT_READ    = 1
    WANT_WRITE   = 2
    AFTER_CREATE = 4
)

// flags: "setmask" field in "CUTOMA_FUSE_SETATTR"
// SET_GOAL_FLAG,SET_DELETE_FLAG are no longer supported
// SET_LENGTH_FLAG,SET_OPENED_FLAG are deprecated
// instead of using FUSE_SETATTR with SET_GOAL_FLAG use FUSE_SETGOAL command
// instead of using FUSE_SETATTR with SET_GOAL_FLAG use FUSE_SETTRASH_TIMEOUT command
// instead of using FUSE_SETATTR with SET_LENGTH_FLAG/SET_OPENED_FLAG use FUSE_TRUNCATE command
const (
    SET_GOAL_FLAG = 1 << iota
    SET_MODE_FLAG
    SET_UID_FLAG
    SET_GID_FLAG
    SET_LENGTH_FLAG
    SET_MTIME_FLAG
    SET_ATIME_FLAG
    SET_OPENED_FLAG
    SET_DELETE_FLAG
)
const ANTOAN_NOP = 0

// CHUNKSERVER <-> CLIENT/CHUNKSERVER
const (
    CUTOCS_READ = 200
    // chunkid:64 version:32 offset:32 size:32
    CSTOCU_READ_STATUS = 201
    // chunkid:64 status:8
    CSTOCU_READ_DATA = 202
    // chunkid:64 blocknum:16 offset:16 size:32 crc:32 size*[ databyte:8 ]

    CUTOCS_WRITE = 210
    // chunkid:64 version:32 N*[ ip:32 port:16 ]
    CSTOCU_WRITE_STATUS = 211
    // chunkid:64 writeid:32 status:8
    CUTOCS_WRITE_DATA = 212
    // chunkid:64 writeid:32 blocknum:16 offset:16 size:32 crc:32 size*[ databyte:8 ]
    CUTOCS_WRITE_FINISH = 213
    // chunkid:64 version:32
)

//ANY <-> CHUNKSERVER
const (
    ANTOCS_CHUNK_CHECKSUM = 300
    // chunkid:64 version:32
    CSTOAN_CHUNK_CHECKSUM = 301
    // chunkid:64 version:32 checksum:32
    // chunkid:64 version:32 status:8

    ANTOCS_CHUNK_CHECKSUM_TAB = 302
    // chunkid:64 version:32
    CSTOAN_CHUNK_CHECKSUM_TAB = 303
    // chunkid:64 version:32 1024*[checksum:32]
    // chunkid:64 version:32 status:8
)

// CLIENT <-> MASTER

// old attr record:
//   type:8 flags:8 mode:16 uid:32 gid:32 atime:32 mtime:32 ctime:32 length:64
//   total: 32B (1+1+2+4+4+4+4+4+8)
//
//   flags: ---DGGGG
//             |\--/
//             |  \------ goal
//             \--------- delete imediatelly

// new attr record:
//   type:8 mode:16 uid:32 gid:32 atime:32 mtime:32 ctime:32 nlink:32 length:64
//   total: 35B
//
//   mode: FFFFMMMMMMMMMMMM
//         \--/\----------/
//           \       \------- mode
//            \-------------- flags
//
//   in case of BLOCKDEV and CHARDEV instead of 'length:64' on the end there is 'mojor:16 minor:16 empty:32'

// NAME type:
// ( leng:8 data:lengB )


const (
    FUSE_REGISTER_BLOB_NOACL = "kFh9mdZsR84l5e675v8bi54VfXaXSYozaU3DSz9AsLLtOtKipzb9aQNkxeOISx64"
    // CUTOMA:
    //  clientid:32 [ version:32 ]
    // MATOCU:
    //  clientid:32
    //  status:8

    FUSE_REGISTER_BLOB_TOOLS_NOACL = "kFh9mdZsR84l5e675v8bi54VfXaXSYozaU3DSz9AsLLtOtKipzb9aQNkxeOISx63"
    // CUTOMA:
    //  -
    // MATOCU:
    //  status:8

    FUSE_REGISTER_BLOB_ACL = "DjI1GAQDULI5d2YjA26ypc3ovkhjvhciTQVx3CS4nYgtBoUcsljiVpsErJENHaw0"

    REGISTER_GETRANDOM = uint8(1)
    // rcode==1: generate random blob
    // CUTOMA:
    //  rcode:8
    // MATOCU:
    //  randomblob:32B

    REGISTER_NEWSESSION = uint8(2)
    // rcode==2: first register
    // CUTOMA:
    //  rcode:8 version:32 ileng:32 info:ilengB pleng:32 path:plengB [ passcode:16B ]
    // MATOCU:
    //  sessionid:32 sesflags:8 rootuid:32 rootgid:32
    //  status:8

    REGISTER_RECONNECT = uint8(3)
    // rcode==3: mount reconnect
    // CUTOMA:
    //  rcode:8 sessionid:32 version:32
    // MATOCU:
    //  status:8

    REGISTER_TOOLS = uint8(4)
    // rcode==4: tools connect
    // CUTOMA:
    //  rcode:8 sessionid:32 version:32
    // MATOCU:
    //  status:8

    REGISTER_NEWMETASESSION = uint8(5)
    // rcode==5: first register
    // CUTOMA:
    //  rcode:8 version:32 ileng:32 info:ilengB [ passcode:16B ]
    // MATOCU:
    //  sessionid:32 sesflags:8
    //  status:8

    CUTOMA_FUSE_REGISTER = 400
    // blob:64B ... (depends on blob - see blob descriptions above)
    MATOCU_FUSE_REGISTER = 401
    // depends on blob - see blob descriptions above
    CUTOMA_FUSE_STATFS = 402
    // msgid:32 -
    MATOCU_FUSE_STATFS = 403
    // msgid:32 totalspace:64 availspace:64 trashspace:64 inodes:32
    CUTOMA_FUSE_ACCESS = 404
    // msgid:32 inode:32 uid:32 gid:32 modemask:8
    MATOCU_FUSE_ACCESS = 405
    // msgid:32 status:8
    CUTOMA_FUSE_LOOKUP = 406
    // msgid:32 inode:32 name:NAME uid:32 gid:32
    MATOCU_FUSE_LOOKUP = 407
    // msgid:32 status:8
    // msgid:32 inode:32 attr:35B
    CUTOMA_FUSE_GETATTR = 408
    // msgid:32 inode:32
    // msgid:32 inode:32 uid:32 gid:32
    MATOCU_FUSE_GETATTR = 409
    // msgid:32 status:8
    // msgid:32 attr:35B
    CUTOMA_FUSE_SETATTR = 410
    // msgid:32 inode:32 uid:32 gid:32 setmask:8 attr:32B   - compatibility with very old version
    // msgid:32 inode:32 uid:32 gid:32 setmask:16 attr:32B  - compatibility with old version
    // msgid:32 inode:32 uid:32 gid:32 setmask:8 attrmode:16 attruid:32 attrgid:32 attratime:32 attrmtime:32
    MATOCU_FUSE_SETATTR = 411
    // msgid:32 status:8
    // msgid:32 attr:35B
    CUTOMA_FUSE_READLINK = 412
    // msgid:32 inode:32
    MATOCU_FUSE_READLINK = 413
    // msgid:32 status:8
    // msgid:32 length:32 path:lengthB
    CUTOMA_FUSE_SYMLINK = 414
    // msgid:32 inode:32 name:NAME length:32 path:lengthB uid:32 gid:32
    MATOCU_FUSE_SYMLINK = 415
    // msgid:32 status:8
    // msgid:32 inode:32 attr:35B
    CUTOMA_FUSE_MKNOD = 416
    // msgid:32 inode:32 name:NAME type:8 mode:16 uid:32 gid:32 rdev:32
    MATOCU_FUSE_MKNOD = 417
    // msgid:32 status:8
    // msgid:32 inode:32 attr:35B
    CUTOMA_FUSE_MKDIR = 418
    // msgid:32 inode:32 name:NAME mode:16 uid:32 gid:32
    MATOCU_FUSE_MKDIR = 419
    // msgid:32 status:8
    // msgid:32 inode:32 attr:35B
    CUTOMA_FUSE_UNLINK = 420
    // msgid:32 inode:32 name:NAME uid:32 gid:32
    MATOCU_FUSE_UNLINK = 421
    // msgid:32 status:8
    CUTOMA_FUSE_RMDIR = 422
    // msgid:32 inode:32 name:NAME uid:32 gid:32
    MATOCU_FUSE_RMDIR = 423
    // msgid:32 status:8
    CUTOMA_FUSE_RENAME = 424
    // msgid:32 inode_src:32 name_src:NAME inode_dst:32 name_dst:NAME uid:32 gid:32
    MATOCU_FUSE_RENAME = 425
    // msgid:32 status:8
    CUTOMA_FUSE_LINK = 426
    // msgid:32 inode:32 inode_dst:32 name_dst:NAME uid:32 gid:32
    MATOCU_FUSE_LINK = 427
    // msgid:32 status:8
    // msgid:32 inode:32 attr:35B
    CUTOMA_FUSE_GETDIR = 428
    // msgid:32 inode:32 uid:32 gid:32 - old version (works like new version with flags==0)
    // msgid:32 inode:32 uid:32 gid:32 flags:8
    MATOCU_FUSE_GETDIR = 429
    // msgid:32 status:8
    // msgid:32 N*[ name:NAME inode:32 type:8 ] - when GETDIR_FLAG_WITHATTR in flags is not set
    // msgid:32 N*[ name:NAME inode:32 type:35B ]   - when GETDIR_FLAG_WITHATTR in flags is set
    CUTOMA_FUSE_OPEN = 430
    // msgid:32 inode:32 uid:32 gid:32 flags:8
    MATOCU_FUSE_OPEN = 431
    // msgid:32 status:8
    // since 1.6.9 if no error:
    // msgid:32 attr:35B

    CUTOMA_FUSE_READ_CHUNK = 432
    // msgid:32 inode:32 chunkindx:32
    MATOCU_FUSE_READ_CHUNK = 433
    // msgid:32 status:8
    // msgid:32 length:64 chunkid:64 version:32 N*[ip:32 port:16]
    // msgid:32 length:64 srcs:8 srcs*[chunkid:64 version:32 ip:32 port:16] - not implemented
    CUTOMA_FUSE_WRITE_CHUNK = 434 /* it creates, duplicates or sets new version of chunk if necessary */
    // msgid:32 inode:32 chunkindx:32
    MATOCU_FUSE_WRITE_CHUNK = 435
    // msgid:32 status:8
    // msgid:32 length:64 chunkid:64 version:32 N*[ip:32 port:16]
    CUTOMA_FUSE_WRITE_CHUNK_END = 436
    // msgid:32 chunkid:64 inode:32 length:64
    MATOCU_FUSE_WRITE_CHUNK_END = 437
    // msgid:32 status:8


    CUTOMA_FUSE_APPEND = 438
    // msgid:32 inode:32 srcinode:32 uid:32 gid:32 - append to existing element
    MATOCU_FUSE_APPEND = 439
    // msgid:32 status:8


    CUTOMA_FUSE_CHECK = 440
    // msgid:32 inode:32
    MATOCU_FUSE_CHECK = 441
    // msgid:32 status:8
    // msgid:32 N*[ copies:8 chunks:16 ]

    CUTOMA_FUSE_GETTRASHTIME = 442
    // msgid:32 inode:32 gmode:8
    MATOCU_FUSE_GETTRASHTIME = 443
    // msgid:32 status:8
    // msgid:32 tdirs:32 tfiles:32 tdirs*[ trashtime:32 dirs:32 ] tfiles*[ trashtime:32 files:32 ]
    CUTOMA_FUSE_SETTRASHTIME = 444
    // msgid:32 inode:32 uid:32 trashtimeout:32 smode:8
    MATOCU_FUSE_SETTRASHTIME = 445
    // msgid:32 status:8
    // msgid:32 changed:32 notchanged:32 notpermitted:32

    CUTOMA_FUSE_GETGOAL = 446
    // msgid:32 inode:32 gmode:8
    MATOCU_FUSE_GETGOAL = 447
    // msgid:32 status:8
    // msgid:32 gdirs:8 gfiles:8 gdirs*[ goal:8 dirs:32 ] gfiles*[ goal:8 files:32 ]

    CUTOMA_FUSE_SETGOAL = 448
    // msgid:32 inode:32 uid:32 goal:8 smode:8
    MATOCU_FUSE_SETGOAL = 449
    // msgid:32 status:8
    // msgid:32 changed:32 notchanged:32 notpermitted:32

    CUTOMA_FUSE_GETTRASH = 450
    // msgid:32
    MATOCU_FUSE_GETTRASH = 451
    // msgid:32 status:8
    // msgid:32 N*[ name:NAME inode:32 ]

    CUTOMA_FUSE_GETDETACHEDATTR = 452
    // msgid:32 inode:32 dtype:8
    MATOCU_FUSE_GETDETACHEDATTR = 453
    // msgid:32 status:8
    // msgid:32 attr:35B

    CUTOMA_FUSE_GETTRASHPATH = 454
    // msgid:32 inode:32
    MATOCU_FUSE_GETTRASHPATH = 455
    // msgid:32 status:8
    // msgid:32 length:32 path:lengthB

    CUTOMA_FUSE_SETTRASHPATH = 456
    // msgid:32 inode:32 length:32 path:lengthB
    MATOCU_FUSE_SETTRASHPATH = 457
    // msgid:32 status:8

    CUTOMA_FUSE_UNDEL = 458
    // msgid:32 inode:32
    MATOCU_FUSE_UNDEL = 459
    // msgid:32 status:8
    CUTOMA_FUSE_PURGE = 460
    // msgid:32 inode:32
    MATOCU_FUSE_PURGE = 461
    // msgid:32 status:8

    CUTOMA_FUSE_GETDIRSTATS = 462
    // msgid:32 inode:32
    MATOCU_FUSE_GETDIRSTATS = 463
    // msgid:32 status:8
    // msgid:32 inodes:32 dirs:32 files:32 ugfiles:32 mfiles:32 chunks:32 ugchunks:32 mchunks32 length:64 size:64 gsize:64

    CUTOMA_FUSE_TRUNCATE = 464
    // msgid:32 inode:32 [opened:8] uid:32 gid:32 opened:8 length:64
    MATOCU_FUSE_TRUNCATE = 465
    // msgid:32 status:8
    // msgid:32 attr:35B

    CUTOMA_FUSE_REPAIR = 466
    // msgid:32 inode:32 uid:32 gid:32
    MATOCU_FUSE_REPAIR = 467
    // msgid:32 status:8
    // msgid:32 notchanged:32 erased:32 repaired:32

    CUTOMA_FUSE_SNAPSHOT = 468
    // msgid:32 inode:32 inode_dst:32 name_dst:NAME uid:32 gid:32 canoverwrite:8
    MATOCU_FUSE_SNAPSHOT = 469
    // msgid:32 status:8

    CUTOMA_FUSE_GETRESERVED = 470
    // msgid:32
    MATOCU_FUSE_GETRESERVED = 471
    // msgid:32 status:8
    // msgid:32 N*[ name:NAME inode:32 ]

    CUTOMA_FUSE_GETEATTR = 472
    // msgid:32 inode:32 gmode:8
    MATOCU_FUSE_GETEATTR = 473
    // msgid:32 status:8
    // msgid:32 eattrdirs:8 eattrfiles:8 eattrdirs*[ eattr:8 dirs:32 ] eattrfiles*[ eattr:8 files:32 ]

    CUTOMA_FUSE_SETEATTR = 474
    // msgid:32 inode:32 uid:32 eattr:8 smode:8
    MATOCU_FUSE_SETEATTR = 475
    // msgid:32 status:8
    // msgid:32 changed:32 notchanged:32 notpermitted:32

    CUTOMA_FUSE_QUOTACONTROL = 476
    // msgid:32 inode:32 qflags:8 - delete quota
    // msgid:32 inode:32 qflags:8 sinodes:32 slength:64 ssize:64 srealsize:64 hinodes:32 hlength:64 hsize:64 hrealsize:64 - set quota
    MATOCU_FUSE_QUOTACONTROL = 477
    // msgid:32 status:8
    // msgid:32 qflags:8 sinodes:32 slength:64 ssize:64 srealsize:64 hinodes:32 hlength:64 hsize:64 hrealsize:64 curinodes:32 curlength:64 cursize:64 currealsize:64

    // special - reserved (opened) inodes - keep opened files.
    CUTOMA_FUSE_RESERVED_INODES = 499
    // N*[inode:32]
)


var errtab = []string{
    "OK",
    "Operation not permitted",
    "Not a directory",
    "No such file or directory",
    "Permission denied",
    "File exists",
    "Invalid argument",
    "Directory not empty",
    "Chunk lost",
    "Out of memory",
    "Index too big",
    "Chunk locked",
    "No chunk servers",
    "No such chunk",
    "Chunk is busy",
    "Incorrect register BLOB",
    "None of chunk servers performed requested operation",
    "File not opened",
    "Write not started",
    "Wrong chunk version",
    "Chunk already exists",
    "No space left",
    "IO error",
    "Incorrect block number",
    "Incorrect size",
    "Incorrect offset",
    "Can't connect",
    "Incorrect chunk id",
    "Disconnected",
    "CRC error",
    "Operation delayed",
    "Can't create path",
    "Data mismatch",
    "Read-only file system",
    "Quota exceeded",
    "Bad session id",
    "Password is needed",
    "Incorrect password",
    "Unknown MFS error",
}

func mfs_strerror(code int) string {
    if code > ERROR_MAX {
        code = ERROR_MAX
    }
    return errtab[code]
}

const (
    S_IFMT   = 0170000 /* type of file */
    S_IFIFO  = 0010000 /* named pipe (fifo) */
    S_IFCHR  = 0020000 /* character special */
    S_IFDIR  = 0040000 /* directory */
    S_IFBLK  = 0060000 /* block special */
    S_IFREG  = 0100000 /* regular */
    S_IFLNK  = 0120000 /* symbolic link */
    S_IFSOCK = 0140000 /* socket */
    S_IFWHT  = 0160000 /* whiteout */
    S_ISUID  = 0004000 /* set user id on execution */
    S_ISGID  = 0002000 /* set group id on execution */
    S_ISVTX  = 0001000 /* save swapped text even after use */
    S_IRUSR  = 0000400 /* read permission, owner */
    S_IWUSR  = 0000200 /* write permission, owner */
    S_IXUSR  = 0000100 /* execute/search permission, owner */
)
