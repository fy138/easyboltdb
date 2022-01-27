package easyboltdb

import (
	"github.com/boltdb/bolt"
)

type EasyBolt struct {
	DB     *bolt.DB
	Bucket string
}

func NewEasyBolt(dbname, bucket string) (*EasyBolt, error) {
	db, err := bolt.Open(dbname, 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		//判断要创建的表是否存在
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			//创建表
			_, err := tx.CreateBucket([]byte(bucket))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sdb := &EasyBolt{}
	sdb.DB = db
	sdb.Bucket = bucket
	return sdb, nil
}
func (s *EasyBolt) Update(key, value string) error {
	err := s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.Bucket))
		if b != nil {
			err := b.Put([]byte(key), []byte(value))
			if err != nil {
				return err
			}
		}
		return nil
	})

	//更新数据库失败
	if err != nil {
		return err
	}
	return nil
}
func (s *EasyBolt) View(key string) (str string, err error) {
	err = s.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.Bucket))
		if b != nil {
			data := b.Get([]byte(key))
			str = string(data)
			return nil
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return
}
func (s *EasyBolt) Close() {
	s.DB.Close()
}
