package nvwa_milvus

import (
	"Nvwa/logger"
	"context"
	"github.com/milvus-io/milvus-proto/go-api/commonpb"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"strconv"
)

var milvusClient client.Client

var (
	memoryVectorCollection = "memory_vector"
	dimension              = 1536 // 维度
	indexFileSize          = 1024
	emptyExpression        = ""
)

func InitMilvus(address string) {
	var err error
	milvusClient, err = client.NewGrpcClient(context.Background(), address)
	if err != nil {
		logger.NvwaLog.Errorf("初始化连接Milvus错误，error=%s", err.Error())
		return
	}
	hasMemoryCollection, err := milvusClient.HasCollection(context.Background(), memoryVectorCollection)
	if err != nil {
		logger.NvwaLog.Errorf("HasCollection %s 报错，error=%s", memoryVectorCollection, err)
		return
	}
	if hasMemoryCollection {
		// 清空memory集合
		err = milvusClient.DropCollection(context.Background(), memoryVectorCollection)
		if err != nil {
			logger.NvwaLog.Errorf("DropCollection %s 报错，error=%s", memoryVectorCollection, err)
			return
		}
	}
	// 创建初始集合
	collectionSchema := &entity.Schema{
		CollectionName: memoryVectorCollection,
		Description:    "agent memory",
		AutoID:         false,
		Fields:         memoryFields,
	}
	err = milvusClient.CreateCollection(context.Background(), collectionSchema, 2)
	if err != nil {
		logger.NvwaLog.Errorf("Milvus创建memory集合错误，error=%s", err.Error())
		return
	}
	// 为embedding创建索引
	index, err := entity.NewIndexIvfFlat(entity.L2, 128)
	if err != nil {
		logger.NvwaLog.Errorf("Milvus新建索引错误，error=%s", err.Error())
		return
	}
	err = milvusClient.CreateIndex(context.Background(), memoryVectorCollection, embeddingColumnName, index, false)
	if err != nil {
		logger.NvwaLog.Errorf("Milvus创建memory集合中embedding索引错误，error=%s", err.Error())
		return
	}
}

// InsertAgentMemoryEmbedding 存储agent记忆文本潜入后的向量
func InsertAgentMemoryEmbedding(memoryId int64, memoryEmbedding []float32) {
	// 根据agentId选择或创建分区
	//partition := strconv.FormatInt(agentId, 10)
	//hasPartition, err := milvusClient.HasPartition(context.Background(), memoryVectorCollection, partition)
	//if err != nil {
	//	logger.NvwaLog.Errorf("HasPartition %s-%d 报错，error=%s", memoryVectorCollection, agentId, err)
	//	return
	//}
	//if !hasPartition {
	//	err = milvusClient.CreatePartition(context.Background(), memoryVectorCollection, partition)
	//	if err != nil {
	//		logger.NvwaLog.Errorf("CreatePartition %s-%d 报错，error=%s", memoryVectorCollection, agentId, err)
	//		return
	//	}
	//}
	idCol := entity.NewColumnInt64(embeddingColumnName, []int64{memoryId})
	embeddingCol := entity.NewColumnFloatVector(embeddingColumnName, dimension, [][]float32{memoryEmbedding})
	idColumn, err := milvusClient.Insert(context.Background(), memoryVectorCollection, "", idCol, embeddingCol)
	if err != nil {
		logger.NvwaLog.Errorf("Insert %s-%d 报错，error=%s", memoryVectorCollection, memoryId, err)
		return
	}
	logger.NvwaLog.Infof("向memory collection中插入一条，id=%v", idColumn)
}

func SearchAgentSimilarityMemory(agentId int, search []float32, topK int) []string {
	// 根据agentId选择或创建分区
	partition := strconv.Itoa(agentId)
	// search 前必须先load到内存
	loadState, err := milvusClient.GetLoadState(context.Background(), memoryVectorCollection, []string{partition})
	if err != nil {
		logger.NvwaLog.Errorf("SearchAgentSimilarityMemory %s-%d GetLoadState报错，error=%s", memoryVectorCollection, agentId, err)
		return nil
	}
	if loadState == entity.LoadStateNotLoad {
		var indexState entity.IndexState
		indexState, err = milvusClient.GetIndexState(context.Background(), memoryVectorCollection, embeddingColumnName)
		if indexState == entity.IndexState(commonpb.IndexState_IndexStateNone) {
			logger.NvwaLog.Errorf("SearchAgentSimilarityMemory %s-%d GetIndexState报错，error=%ve", memoryVectorCollection, agentId, err)
		}
		// LoadPartitions前必须先创建索引index
		err = milvusClient.LoadPartitions(context.Background(), memoryVectorCollection, []string{partition}, true)
		if err != nil {
			logger.NvwaLog.Errorf("SearchAgentSimilarityMemory %s-%d LoadPartitions报错，error=%v", memoryVectorCollection, agentId, err)
			return nil
		}
	}
	searchVector := []entity.Vector{entity.FloatVector(search)}
	searchParam, err := entity.NewIndexIvfFlatSearchParam(16)
	searchResult, err := milvusClient.Search(context.Background(), memoryVectorCollection, []string{partition}, emptyExpression, []string{contentColumnName}, searchVector, embeddingColumnName, entity.L2, topK, searchParam)
	if err != nil {
		logger.NvwaLog.Errorf("SearchAgentSimilarityMemory %s-%d 报错，error=%s", memoryVectorCollection, agentId, err)
		return nil
	}
	var result []string
	for _, sr := range searchResult {
		for _, field := range sr.Fields {
			if field.Type() == entity.FieldTypeVarChar {
				f := field.(*entity.ColumnVarChar)
				result = append(result, f.Data()...)
			}
		}
	}
	// 查询完释放内存
	err = milvusClient.DropPartition(context.Background(), memoryVectorCollection, partition)
	if err == nil {
		logger.NvwaLog.Errorf("DropPartition %s error:%v", partition, err)
	}
	return result
}
